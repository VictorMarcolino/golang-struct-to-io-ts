package generators

import (
	"fmt"
	"go/constant"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"reflect"
	"sort"
	"strings"
)

// IsEnumType checks if the given reflect.Type has any matching constants in its package.
func IsEnumType(t reflect.Type) bool {
	return len(GetEnumConstantsAsMap(t)) > 0
}

// GetEnumConstantsAsMap extracts enum constants associated with the given type
// and returns a map of "constant name" -> "constant value"
func GetEnumConstantsAsMap(t reflect.Type) map[string]interface{} {
	if t == nil || t.PkgPath() == "" || t.Name() == "" {
		return nil
	}
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo,
	}
	pkgs, err := packages.Load(cfg, t.PkgPath())
	if err != nil {
		log.Printf("Error loading package: %v\n", err)
		return nil
	}
	if packages.PrintErrors(pkgs) > 0 {
		return nil
	}
	results := make(map[string]interface{})
	for _, pkg := range pkgs {
		scope := pkg.Types.Scope()
		for _, name := range scope.Names() {
			obj := scope.Lookup(name)
			if c, ok := obj.(*types.Const); ok {
				if named, ok := c.Type().(*types.Named); ok && named.Obj().Name() == t.Name() {
					val := c.Val() // constant.Value

					switch val.Kind() {
					case constant.String:
						// For string constants
						results[c.Name()] = constant.StringVal(val)
					case constant.Bool:
						results[c.Name()] = constant.BoolVal(val)
					case constant.Int:
						// For integer constants (returning int64 if it fits)
						i64, _ := constant.Int64Val(val)
						results[c.Name()] = i64
					case constant.Float:
						// For float constants
						f, _ := constant.Float64Val(val)
						results[c.Name()] = f
					default:
						// Fallback: store exact string
						results[c.Name()] = val.ExactString()
					}
				}
			}
		}
	}
	return results
}

func GetIoTsEnumText(t reflect.Type) string {
	constants := GetEnumConstantsAsMap(t)
	if len(constants) == 0 {
		panic("No constants found for type")
	}

	// Sort the constant names for stable output (optional).
	names := make([]string, 0, len(constants))
	for name := range constants {
		names = append(names, name)
	}
	sort.Strings(names)

	constLines := make([]string, 0, len(constants))
	literalLines := make([]string, 0, len(constants))

	for _, name := range names {
		value := constants[name]

		// Distinguish string vs. numeric
		switch v := value.(type) {
		case string:
			// Strings get quoted
			constLines = append(constLines,
				fmt.Sprintf(`export const %s%s = "%s" as const;`, t.Name(), name, v),
			)
			literalLines = append(literalLines,
				fmt.Sprintf(`t.literal(%s%s)`, t.Name(), name),
			)

		case int64, int, float64:
			// Numeric constants: no quotes
			constLines = append(constLines,
				fmt.Sprintf(`export const %s%s = %v as const;`, t.Name(), name, v),
			)
			literalLines = append(literalLines,
				fmt.Sprintf(`t.literal(%s%s)`, t.Name(), name),
			)

		default:
			// Fallback to string representation, if needed
			constVal := fmt.Sprintf("%v", v)
			constLines = append(constLines,
				fmt.Sprintf(`export const %s%s = "%s" as const;`, t.Name(), name, constVal),
			)
			literalLines = append(literalLines,
				fmt.Sprintf(`t.literal(%s%s)`, t.Name(), name),
			)
		}
	}

	allConsts := strings.Join(constLines, "\n")
	allLiterals := strings.Join(literalLines, ",\n")

	typeDef := fmt.Sprintf(`%s

export const %sC = t.union([
%s
]);

export type %s = t.TypeOf<typeof %sC>;

`, allConsts, t.Name(), allLiterals, t.Name(), t.Name())

	return typeDef
}
