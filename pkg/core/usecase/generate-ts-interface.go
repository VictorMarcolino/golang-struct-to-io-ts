package usecase

import (
	"fmt"
	"reflect"
	"strings"
)

// TypeScriptGeneratorOptions defines options for generating TypeScript interfaces
type TypeScriptGeneratorOptions struct {
	TreatArraysAsOptional bool
}

// TypeScriptGenerator encapsulates the logic to generate TypeScript interfaces
type TypeScriptGenerator struct {
	generatedStructs map[string]bool
	options          TypeScriptGeneratorOptions
}

// NewTypeScriptGenerator creates a new instance of TypeScriptGenerator with the provided options
func NewTypeScriptGenerator(options ...TypeScriptGeneratorOptions) *TypeScriptGenerator {
	chosenOptions := TypeScriptGeneratorOptions{
		TreatArraysAsOptional: false,
	}
	if len(options) != 0 {
		chosenOptions = options[0]
	}
	return &TypeScriptGenerator{
		generatedStructs: make(map[string]bool),
		options:          chosenOptions,
	}
}

// Generate takes any struct and generates its corresponding TypeScript interface
func (g *TypeScriptGenerator) Generate(inputStruct interface{}) (string, error) {
	t := reflect.TypeOf(inputStruct)

	// Ensure the input is a struct
	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("input is not a struct")
	}

	var interfaces strings.Builder
	g.processStruct(t, &interfaces)
	return interfaces.String(), nil
}

// processStruct processes a struct and generates its TypeScript interface
func (g *TypeScriptGenerator) processStruct(t reflect.Type, interfaces *strings.Builder) {
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

	// Then generate the TypeScript interface for the current struct
	interfaces.WriteString(fmt.Sprintf("interface %s {\n", t.Name()))

	// Process each field of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if jsonTag := field.Tag.Get("json"); jsonTag != "" && jsonTag != "-" {
			g.processField(field, interfaces)
		}
	}

	interfaces.WriteString("}\n\n")
}

// processField processes a single field and adds it to the TypeScript interface
func (g *TypeScriptGenerator) processField(field reflect.StructField, interfaces *strings.Builder) {
	jsonTag := field.Tag.Get("json")
	isOptional := g.isFieldOptional(field) // Only pointers should be optional in TypeScript
	tsType := goTypeToTSType(field.Type, isOptional)

	// Write the field to the TypeScript interface
	if isOptional {
		interfaces.WriteString(fmt.Sprintf("  %s?: %s;\n", jsonTag, tsType))
	} else {
		interfaces.WriteString(fmt.Sprintf("  %s: %s;\n", jsonTag, tsType))
	}

	// Recursively process nested structs
	if isStructType(field.Type) {
		g.processStruct(g.dereferenceType(field.Type), interfaces)
	}
}

// isStructProcessed checks if a struct has already been processed
func (g *TypeScriptGenerator) isStructProcessed(t reflect.Type) bool {
	return g.generatedStructs[t.Name()]
}

// markStructProcessed marks a struct as processed
func (g *TypeScriptGenerator) markStructProcessed(t reflect.Type) {
	g.generatedStructs[t.Name()] = true
}

// dereferenceType dereferences pointer types to get the underlying type
func (g *TypeScriptGenerator) dereferenceType(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}

// goTypeToTSType converts a Go type to its corresponding TypeScript type
func goTypeToTSType(goType reflect.Type, isOptional bool) string {
	switch goType.Kind() {
	case reflect.String:
		return "string"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float64, reflect.Float32:
		return "number"
	case reflect.Bool:
		return "boolean"
	case reflect.Slice, reflect.Array:
		return goTypeToTSType(goType.Elem(), false) + "[]"
	case reflect.Ptr:
		return goTypeToTSType(goType.Elem(), true)
	case reflect.Struct:
		return goType.Name()
	default:
		return "any"
	}
}

// isFieldOptional determines if a field should be optional in TypeScript (if it's a pointer)
func (g *TypeScriptGenerator) isFieldOptional(field reflect.StructField) bool {
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

// isStructType checks if the given type is a struct or a pointer to a struct
func isStructType(t reflect.Type) bool {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Kind() == reflect.Struct
}
