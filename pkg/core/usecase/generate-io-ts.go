package usecase

import (
	"fmt"
	"reflect"
	"strings"
)

// IoTsGenerator encapsulates the logic to generate io-ts types
type IoTsGenerator struct {
	generatedStructs map[string]bool
	options          TypeScriptGeneratorOptions
}

// NewIoTsGenerator creates a new instance of IoTsGenerator with the provided options
func NewIoTsGenerator(options ...TypeScriptGeneratorOptions) *IoTsGenerator {
	chosenOptions := TypeScriptGeneratorOptions{
		TreatArraysAsOptional: false,
	}
	if len(options) != 0 {
		chosenOptions = options[0]
	}
	return &IoTsGenerator{
		generatedStructs: make(map[string]bool),
		options:          chosenOptions,
	}
}

// Generate takes any struct and generates its corresponding io-ts type
func (g *IoTsGenerator) Generate(inputStruct interface{}) (string, error) {
	t := reflect.TypeOf(inputStruct)

	// Ensure the input is a struct
	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("input is not a struct")
	}

	var interfaces strings.Builder
	g.processStruct(t, &interfaces)
	return interfaces.String(), nil
}

// processStruct processes a struct and generates its io-ts type
func (g *IoTsGenerator) processStruct(t reflect.Type, interfaces *strings.Builder) {
	// Dereference pointer types
	t = g.dereferenceType(t)

	// Avoid reprocessing the same struct
	if g.isStructProcessed(t) {
		return
	}

	// First, recursively process any nested structs or elements of a slice
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldType := g.dereferenceType(field.Type)

		// If the field is a slice, process the element type (recursively handle nested structs)
		if fieldType.Kind() == reflect.Slice {
			elementType := g.dereferenceType(fieldType.Elem())
			if isStructType(elementType) {
				g.processStruct(elementType, interfaces) // Process child structs inside slices
			}
		} else if isStructType(fieldType) {
			g.processStruct(fieldType, interfaces) // Process regular child structs
		}
	}

	// Now mark the current struct as processed
	g.markStructProcessed(t)

	// Then generate the io-ts type for the current struct
	interfaces.WriteString(fmt.Sprintf("const %s = t.type({\n", t.Name()))

	// Process each field of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if jsonTag := field.Tag.Get("json"); jsonTag != "" && jsonTag != "-" {
			g.processField(field, interfaces)
		}
	}

	interfaces.WriteString("});\n\n")
}

// processField processes a single field and adds it to the io-ts type
func (g *IoTsGenerator) processField(field reflect.StructField, interfaces *strings.Builder) {
	jsonTag := field.Tag.Get("json")

	// Extract the field name (ignore ",omitempty" or other options)
	jsonFieldName := strings.Split(jsonTag, ",")[0]

	isOptional := g.isFieldOptional(field)
	ioTsType := goTypeToIoTSType(field.Type, isOptional)

	// Handle "omitempty" fields
	if strings.Contains(jsonTag, ",omitempty") {
		isOptional = true
	}

	// Write the field to the io-ts type
	if isOptional {
		interfaces.WriteString(fmt.Sprintf("  %s: t.union([%s, t.undefined]),\n", jsonFieldName, ioTsType))
	} else {
		interfaces.WriteString(fmt.Sprintf("  %s: %s,\n", jsonFieldName, ioTsType))
	}

	// Recursively process nested structs
	if isStructType(field.Type) {
		g.processStruct(g.dereferenceType(field.Type), interfaces)
	}
}

// goTypeToIoTSType converts a Go type to its corresponding io-ts type
func goTypeToIoTSType(goType reflect.Type, isOptional bool) string {
	// Special case for map[string]interface{}
	if goType.Kind() == reflect.Map && goType.Key().Kind() == reflect.String && goType.Elem().Kind() == reflect.Interface {
		return "t.record(t.string, t.unknown)"
	}

	switch goType.Kind() {
	case reflect.String:
		return "t.string"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float64, reflect.Float32:
		return "t.number"
	case reflect.Bool:
		return "t.boolean"
	case reflect.Slice, reflect.Array:
		// Check if the slice's element type is a pointer
		elementType := goType.Elem()
		if elementType.Kind() == reflect.Ptr {
			return fmt.Sprintf("t.array(t.union([%s, t.undefined]))", goTypeToIoTSType(elementType.Elem(), false))
		}
		return fmt.Sprintf("t.array(%s)", goTypeToIoTSType(elementType, false))
	case reflect.Ptr:
		return goTypeToIoTSType(goType.Elem(), true)
	case reflect.Struct:
		return fmt.Sprintf("%s", goType.Name())
	default:
		return "t.unknown"
	}
}

// isFieldOptional determines if a field should be optional in io-ts (if it's a pointer or array)
func (g *IoTsGenerator) isFieldOptional(field reflect.StructField) bool {
	fieldType := field.Type

	// Mark arrays as optional if the option is set
	if (fieldType.Kind() == reflect.Slice || fieldType.Kind() == reflect.Array) && g.options.TreatArraysAsOptional {
		return true
	}

	// Mark pointers as optional
	if fieldType.Kind() == reflect.Ptr {
		return true
	}

	return false
}

// Helper methods for struct processing

func (g *IoTsGenerator) isStructProcessed(t reflect.Type) bool {
	return g.generatedStructs[t.Name()]
}

func (g *IoTsGenerator) markStructProcessed(t reflect.Type) {
	g.generatedStructs[t.Name()] = true
}

func (g *IoTsGenerator) dereferenceType(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}
