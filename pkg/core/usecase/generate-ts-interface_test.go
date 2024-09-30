package usecase_test

import (
	"github.com/VictorMarcolino/golang-struct-to-io-ts/pkg/core/usecase"
	"github.com/VictorMarcolino/golang-struct-to-io-ts/pkg/core/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Simple Cases", func() {
	It("should generate correct TypeScript for int field", func() {
		type SimpleCase struct {
			Age int `json:"age"`
		}
		user := SimpleCase{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)
		expected := `interface SimpleCase { age: number; }`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct TypeScript for string field", func() {
		type SimpleCase struct {
			Name string `json:"name"`
		}
		user := SimpleCase{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)
		expected := `interface SimpleCase { name: string; }`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct TypeScript for bool field", func() {
		type SimpleCase struct {
			IsActive bool `json:"is_active"`
		}
		user := SimpleCase{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)
		expected := `interface SimpleCase { is_active: boolean; }`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct TypeScript for float64 field", func() {
		type SimpleCase struct {
			Price float64 `json:"price"`
		}
		user := SimpleCase{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)
		expected := `interface SimpleCase { price: number; }`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct TypeScript for uint field", func() {
		type SimpleCase struct {
			Count uint `json:"count"`
		}
		user := SimpleCase{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)
		expected := `interface SimpleCase { count: number; }`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct TypeScript for pointer field", func() {
		type SimpleCase struct {
			ZipCode *string `json:"zip_code"`
		}
		user := SimpleCase{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)
		expected := `interface SimpleCase { zip_code?: string; }`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct TypeScript for pointer int field", func() {
		type SimpleCase struct {
			Score *int `json:"score"`
		}
		user := SimpleCase{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)
		expected := `interface SimpleCase { score?: number; }`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct TypeScript for slice field", func() {
		type SimpleCase struct {
			Tags []string `json:"tags"`
		}
		user := SimpleCase{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)
		expected := `interface SimpleCase { tags: string[]; }`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct TypeScript for slice of pointers", func() {
		type SimpleCase struct {
			Scores []*int `json:"scores"`
		}
		user := SimpleCase{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)
		expected := `interface SimpleCase { scores: number[]; }`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})
})

var _ = Describe("Nested Cases", func() {
	It("should generate correct TypeScript for nested struct with int field", func() {
		type NestedCaseChildren struct {
			Age int `json:"age"`
		}
		type NestedCaseFather struct {
			Children NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)

		// The expected output includes both interfaces for the nested child and parent
		expected := `interface NestedCaseChildren { age: number; } interface NestedCaseFather { children: NestedCaseChildren; }`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct TypeScript for nested struct with string field", func() {
		type NestedCaseChildren struct {
			Name string `json:"name"`
		}
		type NestedCaseFather struct {
			Children NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)

		expected := `interface NestedCaseChildren { name: string; } interface NestedCaseFather { children: NestedCaseChildren; }`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct TypeScript for nested struct with bool field", func() {
		type NestedCaseChildren struct {
			IsActive bool `json:"is_active"`
		}
		type NestedCaseFather struct {
			Children NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)

		expected := `interface NestedCaseChildren { is_active: boolean; } interface NestedCaseFather { children: NestedCaseChildren; }`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct TypeScript for nested struct with float64 field", func() {
		type NestedCaseChildren struct {
			Salary float64 `json:"salary"`
		}
		type NestedCaseFather struct {
			Children NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)

		expected := `interface NestedCaseChildren { salary: number; } interface NestedCaseFather { children: NestedCaseChildren; }`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct TypeScript for nested struct with uint field", func() {
		type NestedCaseChildren struct {
			Count uint `json:"count"`
		}
		type NestedCaseFather struct {
			Children NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)

		expected := `interface NestedCaseChildren { count: number; } interface NestedCaseFather { children: NestedCaseChildren; }`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct TypeScript for nested struct with pointer field", func() {
		type NestedCaseChildren struct {
			ZipCode *string `json:"zip_code"`
		}
		type NestedCaseFather struct {
			Children NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)

		expected := `interface NestedCaseChildren { zip_code?: string; } interface NestedCaseFather { children: NestedCaseChildren; }`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct TypeScript for nested struct with slice field", func() {
		type NestedCaseChildren struct {
			Tags []string `json:"tags"`
		}
		type NestedCaseFather struct {
			Children NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)

		expected := `interface NestedCaseChildren { tags: string[]; } interface NestedCaseFather { children: NestedCaseChildren; }`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct TypeScript for nested struct with array of pointers", func() {
		type NestedCaseChildren struct {
			Scores []*int `json:"scores"`
		}
		type NestedCaseFather struct {
			Children NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)

		expected := `interface NestedCaseChildren { scores: number[]; } interface NestedCaseFather { children: NestedCaseChildren; }`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})
	It("should generate correct TypeScript for nested struct with father having array as children", func() {
		type NestedCaseChildren struct {
			Scores []*int `json:"scores"`
		}
		type NestedCaseFather struct {
			Children []NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)

		expected := `interface NestedCaseChildren { scores: number[]; } interface NestedCaseFather { children: NestedCaseChildren[]; }`

		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})
	It("should generate correct TypeScript for nested struct with father having pointer array as children", func() {
		type NestedCaseChildren struct {
			Scores []*int `json:"scores"`
		}
		type NestedCaseFather struct {
			Children []*NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewTypeScriptGenerator()
		result, err := generator.Generate(user)

		expected := `interface NestedCaseChildren { scores: number[]; } interface NestedCaseFather { children: NestedCaseChildren[]; }`

		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})
	It("should generate correct TypeScript for nested struct with father having pointer array as children, optional array enabled", func() {
		type NestedCaseChildren struct {
			Scores []*int `json:"scores"`
		}
		type NestedCaseFather struct {
			Children []*NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewTypeScriptGenerator(usecase.TypeScriptGeneratorOptions{TreatArraysAsOptional: true})
		result, err := generator.Generate(user)

		expected := `interface NestedCaseChildren { scores?: number[]; } interface NestedCaseFather { children?: NestedCaseChildren[]; }`

		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})
})
