package generators

import (
	"fmt"
	"reflect"
	"strings"
)

// TypeScriptGeneratorOptions defines options for generating TypeScript interfaces
type TypeScriptGeneratorOptions struct {
	TreatArraysAsOptional bool
}

// TypeConverter defines an interface for converting Go types to io-ts types
type TypeConverter interface {
	Convert(goType reflect.Type, isOptional bool) string
}

// FieldProcessor defines an interface for processing struct fields
type FieldProcessor interface {
	ProcessField(field reflect.StructField) string
	ProcessInlineField(field reflect.StructField) []string
}

// CodeBuilder handles the assembly of generated code
type CodeBuilder struct {
	imports         []string
	typeDefinitions []string
	processedTypes  map[string]struct{}
}

// NewCodeBuilder creates a new CodeBuilder instance
func NewCodeBuilder() *CodeBuilder {
	return &CodeBuilder{
		imports:         []string{"import * as t from 'io-ts';\n\n"},
		typeDefinitions: []string{},
		processedTypes:  make(map[string]struct{}),
	}
}

// AddTypeDefinition adds a type definition to the builder
func (cb *CodeBuilder) AddTypeDefinition(typeDef string) {
	cb.typeDefinitions = append(cb.typeDefinitions, typeDef)
}

// IsTypeProcessed checks if a type has already been processed
func (cb *CodeBuilder) IsTypeProcessed(typeKey string) bool {
	_, exists := cb.processedTypes[typeKey]
	return exists
}

// MarkTypeProcessed marks a type as processed
func (cb *CodeBuilder) MarkTypeProcessed(typeKey string) {
	cb.processedTypes[typeKey] = struct{}{}
}

// Build assembles the final code output
func (cb *CodeBuilder) Build() string {
	var sb strings.Builder
	for _, imp := range cb.imports {
		sb.WriteString(imp)
	}
	for _, typeDef := range cb.typeDefinitions {
		sb.WriteString(typeDef)
	}
	return sb.String()
}

// DefaultTypeConverter is the default implementation of TypeConverter
type DefaultTypeConverter struct {
	generator *IoTsGenerator
}

// Convert converts a Go type to its corresponding io-ts type
func (tc *DefaultTypeConverter) Convert(goType reflect.Type, isOptional bool) string {
	goType = dereferenceType(goType)

	// Special case for map[string]interface{}
	if goType.Kind() == reflect.Map && goType.Key().Kind() == reflect.String && goType.Elem().Kind() == reflect.Interface {
		ioTsType := "t.record(t.string, t.unknown)"
		return wrapOptional(ioTsType, isOptional)
	}

	var ioTsType string
	if IsEnumType(goType) {
		tc.generator.generateEnumType(goType)
		ioTsType = fmt.Sprintf("%sC", goType.Name())
		return wrapOptional(ioTsType, isOptional)
	}
	switch goType.Kind() {
	case reflect.String:
		ioTsType = "t.string"
		break
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		ioTsType = "t.number"
		break
	case reflect.Bool:
		ioTsType = "t.boolean"
		break
	case reflect.Slice, reflect.Array:
		elementType := goType.Elem()
		// Check if the element is a pointer
		isElementOptional := elementType.Kind() == reflect.Ptr
		elementIoTsType := tc.Convert(elementType, isElementOptional)
		ioTsType = fmt.Sprintf("t.array(%s)", elementIoTsType)
		break
	case reflect.Struct:
		typeName := goType.Name()
		if typeName == "" {
			// Anonymous struct, generate inline type
			inlineType := tc.generator.generateInlineStruct(goType)
			ioTsType = inlineType
		} else {
			tc.generator.processStruct(goType)
			ioTsType = fmt.Sprintf("%sC", typeName)
		}
		break
	default:
		ioTsType = "t.unknown"
	}

	return wrapOptional(ioTsType, isOptional)
}

// IoTsGenerator encapsulates the logic to generate io-ts types
type IoTsGenerator struct {
	options       TypeScriptGeneratorOptions
	typeConverter TypeConverter
	codeBuilder   *CodeBuilder
}

// NewIoTsGenerator creates a new instance of NewIoTsGenerator with the provided options
func NewIoTsGenerator(options ...TypeScriptGeneratorOptions) *IoTsGenerator {
	chosenOptions := TypeScriptGeneratorOptions{}
	if len(options) != 0 {
		chosenOptions = options[0]
	}
	generator := &IoTsGenerator{
		options:     chosenOptions,
		codeBuilder: NewCodeBuilder(),
	}
	generator.typeConverter = &DefaultTypeConverter{generator: generator}
	return generator
}

// Generate takes any struct and generates its corresponding io-ts type
func (g *IoTsGenerator) Generate(inputStruct interface{}) (string, error) {
	t := reflect.TypeOf(inputStruct)
	t = dereferenceType(t)

	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("input is not a struct")
	}

	g.processStruct(t)
	return g.codeBuilder.Build(), nil
}

// processStruct processes a struct and generates its io-ts type
func (g *IoTsGenerator) processStruct(t reflect.Type) {
	t = dereferenceType(t)

	typeKey := getTypeKey(t)
	if g.codeBuilder.IsTypeProcessed(typeKey) || t.Name() == "" {
		return
	}

	g.processNestedStructs(t)
	g.codeBuilder.MarkTypeProcessed(typeKey)
	typeDef := g.generateIoTsType(t)
	g.codeBuilder.AddTypeDefinition(typeDef)
}

// processNestedStructs processes nested structs within a parent struct
func (g *IoTsGenerator) processNestedStructs(t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if g.shouldSkipField(field) {
			continue
		}

		fieldType := dereferenceType(field.Type)

		// Avoid infinite recursion: skip processing if the field type is the same as the parent type
		if getTypeKey(fieldType) == getTypeKey(t) {
			continue
		}

		if isStructType(fieldType) {
			if strings.Contains(field.Tag.Get("json"), ",inline") {
				g.processNestedStructs(fieldType)
			} else if fieldType.Name() == "" {
				// Anonymous struct
				g.typeConverter.Convert(fieldType, g.isFieldOptional(field))
			} else {
				g.processStruct(fieldType)
			}
		} else if fieldType.Kind() == reflect.Slice || fieldType.Kind() == reflect.Array {
			elementType := dereferenceType(fieldType.Elem())
			// Avoid infinite recursion for slices/arrays of the same type
			if getTypeKey(elementType) == getTypeKey(t) {
				continue
			}
			if isStructType(elementType) {
				if elementType.Name() == "" {
					// Anonymous struct
					g.typeConverter.Convert(elementType, false)
				} else {
					g.processStruct(elementType)
				}
			}
		}
	}
}

// generateIoTsType generates the io-ts type for a struct and returns it as a string
func (g *IoTsGenerator) generateIoTsType(t reflect.Type) string {
	// If the struct is recursive (contains a field of its own type), emit a t.recursion wrapper
	if isRecursiveStruct(t) {
		var fieldLines []string
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if g.shouldSkipField(field) {
				continue
			}
			jsonTag := field.Tag.Get("json")
			jsonFieldName := strings.Split(jsonTag, ",")[0]

			// Work with the raw type to preserve pointer/collection info for Self detection
			rawType := field.Type
			deref := dereferenceType(rawType)
			selfKey := getTypeKey(t)

			// Pointer to self => union with t.undefined (optional)
			if rawType.Kind() == reflect.Ptr && getTypeKey(deref) == selfKey {
				fieldLines = append(fieldLines, fmt.Sprintf("      %s: t.union([Self, t.undefined]),", formatPropertyName(jsonFieldName)))
				continue
			}
			// Slice/array cases that reference self
			if rawType.Kind() == reflect.Slice || rawType.Kind() == reflect.Array {
				el := rawType.Elem()
				elDeref := dereferenceType(el)
				if el.Kind() == reflect.Ptr && getTypeKey(elDeref) == selfKey {
					// []*Self -> t.array(t.union([Self, t.undefined]))
					fieldLines = append(fieldLines, fmt.Sprintf("      %s: t.array(t.union([Self, t.undefined])),", formatPropertyName(jsonFieldName)))
					continue
				}
				if getTypeKey(elDeref) == selfKey {
					// []Self -> t.array(Self)
					fieldLines = append(fieldLines, fmt.Sprintf("      %s: t.array(Self),", formatPropertyName(jsonFieldName)))
					continue
				}
			}

			// Inline fields are not expected to be self in recursion test, but handle generically
			if strings.Contains(jsonTag, ",inline") {
				inlineFields := g.processInlineField(field)
				for _, f := range inlineFields {
					fieldLines = append(fieldLines, "      "+strings.TrimSpace(f))
				}
				continue
			}
			// Use normal conversion for other fields
			isOptional := strings.Contains(jsonTag, ",omitempty") || g.isFieldOptional(field)
			ioTsType := g.typeConverter.Convert(field.Type, isOptional)
			fieldLines = append(fieldLines, fmt.Sprintf("      %s: %s,", formatPropertyName(jsonFieldName), ioTsType))
		}

		// Build the recursion block
		typeDef := fmt.Sprintf("export const %sC = t.recursion(\n  '%s',\n  Self =>\n    t.type({\n%s\n    }),\n);\n\n", t.Name(), t.Name(), strings.Join(fieldLines, "\n"))
		typeDef += fmt.Sprintf("export type %s = t.TypeOf<typeof %sC>;\n", t.Name(), t.Name())
		return typeDef
	}

	// Non-recursive: default behavior
	var fields []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if g.shouldSkipField(field) {
			continue
		}
		if strings.Contains(field.Tag.Get("json"), ",inline") {
			inlineFields := g.processInlineField(field)
			fields = append(fields, inlineFields...)
		} else {
			fieldDef := g.processField(field)
			fields = append(fields, fieldDef)
		}
	}

	typeDef := fmt.Sprintf("export const %sC = t.type({\n%s\n});\n", t.Name(), strings.Join(fields, "\n"))
	typeDef += fmt.Sprintf("export type %s = t.TypeOf<typeof %sC>;\n\n", t.Name(), t.Name())
	return typeDef
}

// processField processes a single field and returns its definition
func (g *IoTsGenerator) processField(field reflect.StructField) string {
	jsonTag := field.Tag.Get("json")
	jsonFieldName := strings.Split(jsonTag, ",")[0]

	isOptional := strings.Contains(jsonTag, ",omitempty") || g.isFieldOptional(field)

	fieldType := dereferenceType(field.Type)
	if IsEnumType(fieldType) {
		// Ensure the enum is generated if it hasn't been yet
		g.generateEnumType(fieldType)

		// Reference the generated union type, e.g. `ValidateSeverityC`
		ioTsType := fmt.Sprintf("%sC", fieldType.Name())

		// If optional, wrap in union with t.undefined
		if isOptional {
			ioTsType = fmt.Sprintf("t.union([%s, t.undefined])", ioTsType)
		}
		return fmt.Sprintf("  %s: %s,", jsonFieldName, ioTsType)
	}

	// Otherwise, use the normal conversion logic
	ioTsType := g.typeConverter.Convert(field.Type, isOptional)

	return fmt.Sprintf("  %s: %s,", formatPropertyName(jsonFieldName), ioTsType)
}

// processInlineField processes an inlined field and returns its fields
func (g *IoTsGenerator) processInlineField(field reflect.StructField) []string {
	fieldType := dereferenceType(field.Type)
	var fields []string

	for i := 0; i < fieldType.NumField(); i++ {
		inlineField := fieldType.Field(i)
		if g.shouldSkipField(inlineField) {
			continue
		}
		if strings.Contains(inlineField.Tag.Get("json"), ",inline") {
			inlineFields := g.processInlineField(inlineField)
			fields = append(fields, inlineFields...)
		} else {
			fieldDef := g.processField(inlineField)
			fields = append(fields, fieldDef)
		}
	}

	return fields
}

// generateInlineStruct generates an inline type for anonymous structs
func (g *IoTsGenerator) generateInlineStruct(t reflect.Type) string {
	fields := g.generateInlineStructFields(t)
	return fmt.Sprintf("t.type({\n%s\n})", strings.Join(fields, "\n"))
}

// generateInlineStructFields collects field definitions from a struct, including embedded fields
func (g *IoTsGenerator) generateInlineStructFields(t reflect.Type) []string {
	var fields []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if g.shouldSkipField(field) {
			continue
		}

		if field.Anonymous {
			// Embedded field, include its fields recursively
			embeddedType := dereferenceType(field.Type)
			if isStructType(embeddedType) {
				embeddedFields := g.generateInlineStructFields(embeddedType)
				fields = append(fields, embeddedFields...)
			} else {
				// Not a struct, treat as a regular field
				fieldDef := g.processField(field)
				fields = append(fields, fieldDef)
			}
		} else {
			fieldDef := g.processField(field)
			fields = append(fields, fieldDef)
		}
	}

	return fields
}

// isFieldOptional determines if a field should be optional in io-ts
func (g *IoTsGenerator) isFieldOptional(field reflect.StructField) bool {
	fieldType := field.Type

	if (fieldType.Kind() == reflect.Slice || fieldType.Kind() == reflect.Array) && g.options.TreatArraysAsOptional {
		return true
	}

	return fieldType.Kind() == reflect.Ptr
}

// shouldSkipField determines if a field should be skipped
func (g *IoTsGenerator) shouldSkipField(field reflect.StructField) bool {
	// Always include anonymous fields (embedded structs)
	if field.Anonymous {
		return false
	}
	jsonTag := field.Tag.Get("json")
	return jsonTag == "" || jsonTag == "-"
}

// Helper functions

func wrapOptional(ioTsType string, isOptional bool) string {
	if isOptional {
		return fmt.Sprintf("t.union([%s, t.undefined])", ioTsType)
	}
	return ioTsType
}

func getTypeKey(t reflect.Type) string {
	return t.PkgPath() + "." + t.Name()
}

func dereferenceType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

// formatPropertyName returns a valid TypeScript object key for property name.
// If name is a valid identifier, it returns as-is. Otherwise, it single-quotes and escapes it.
func formatPropertyName(name string) string {
	if name == "" {
		return "''"
	}
	isIdent := true
	for i, r := range name {
		if i == 0 {
			if !((r == '_') || (r == '$') || (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
				isIdent = false
				break
			}
		} else {
			if !((r == '_') || (r == '$') || (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')) {
				isIdent = false
				break
			}
		}
	}
	if isIdent {
		return name
	}
	escaped := strings.ReplaceAll(name, "\\", "\\\\")
	escaped = strings.ReplaceAll(escaped, "'", "\\'")
	return fmt.Sprintf("'%s'", escaped)
}

func isStructType(t reflect.Type) bool {
	t = dereferenceType(t)
	return t.Kind() == reflect.Struct
}

// isRecursiveStruct checks whether the struct has a field that refers to its own type (directly or via pointer)
func isRecursiveStruct(t reflect.Type) bool {
	for i := 0; i < t.NumField(); i++ {
		fieldType := dereferenceType(t.Field(i).Type)
		if getTypeKey(fieldType) == getTypeKey(t) {
			return true
		}
		// Check slices/arrays of the same type
		if fieldType.Kind() == reflect.Slice || fieldType.Kind() == reflect.Array {
			elem := dereferenceType(fieldType.Elem())
			if getTypeKey(elem) == getTypeKey(t) {
				return true
			}
		}
	}
	return false
}

// generateEnumType generates the io-ts type for an enum and adds it to the builder.
func (g *IoTsGenerator) generateEnumType(t reflect.Type) {
	if g.codeBuilder.IsTypeProcessed(getTypeKey(t)) {
		return
	}
	// For some reason code is not reaching this point
	iotsText := GetIoTsEnumText(t)
	g.codeBuilder.AddTypeDefinition(iotsText)
	g.codeBuilder.MarkTypeProcessed(getTypeKey(t))
}
