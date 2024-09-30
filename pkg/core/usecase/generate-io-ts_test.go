package usecase_test

import (
	"github.com/VictorMarcolino/golang-struct-to-io-ts/pkg/core/usecase"
	"github.com/VictorMarcolino/golang-struct-to-io-ts/pkg/core/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("IO-TS:Simple Cases", func() {
	It("should generate correct io-ts type for int field", func() {
		type SimpleCase struct {
			Age int `json:"age"`
		}
		user := SimpleCase{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `const SimpleCase = t.type({ age: t.number, });`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for string field", func() {
		type SimpleCase struct {
			Name string `json:"name"`
		}
		user := SimpleCase{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `const SimpleCase = t.type({ name: t.string, });`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for bool field", func() {
		type SimpleCase struct {
			IsActive bool `json:"is_active"`
		}
		user := SimpleCase{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `const SimpleCase = t.type({ is_active: t.boolean, });`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for float64 field", func() {
		type SimpleCase struct {
			Price float64 `json:"price"`
		}
		user := SimpleCase{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `const SimpleCase = t.type({ price: t.number, });`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for uint field", func() {
		type SimpleCase struct {
			Count uint `json:"count"`
		}
		user := SimpleCase{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `const SimpleCase = t.type({ count: t.number, });`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for pointer string field", func() {
		type SimpleCase struct {
			ZipCode *string `json:"zip_code"`
		}
		user := SimpleCase{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `const SimpleCase = t.type({ zip_code: t.union([t.string, t.undefined]), });`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for pointer int field", func() {
		type SimpleCase struct {
			Score *int `json:"score"`
		}
		user := SimpleCase{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `const SimpleCase = t.type({ score: t.union([t.number, t.undefined]), });`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for slice field", func() {
		type SimpleCase struct {
			Tags []string `json:"tags"`
		}
		user := SimpleCase{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `const SimpleCase = t.type({ tags: t.array(t.string), });`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for slice of pointers", func() {
		type SimpleCase struct {
			Scores []*int `json:"scores"`
		}
		user := SimpleCase{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `const SimpleCase = t.type({ scores: t.array(t.union([t.number, t.undefined])), });`

		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})
})

var _ = Describe("IO-TS:Nested Cases", func() {
	It("should generate correct io-ts type for nested struct with int field", func() {
		type NestedCaseChildren struct {
			Age int `json:"age"`
		}
		type NestedCaseFather struct {
			Children NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
const NestedCaseChildren = t.type({
  age: t.number,
});

const NestedCaseFather = t.type({
  children: NestedCaseChildren,
});
`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for nested struct with string field", func() {
		type NestedCaseChildren struct {
			Name string `json:"name"`
		}
		type NestedCaseFather struct {
			Children NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
const NestedCaseChildren = t.type({
  name: t.string,
});

const NestedCaseFather = t.type({
  children: NestedCaseChildren,
});
`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for nested struct with bool field", func() {
		type NestedCaseChildren struct {
			IsActive bool `json:"is_active"`
		}
		type NestedCaseFather struct {
			Children NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
const NestedCaseChildren = t.type({
  is_active: t.boolean,
});

const NestedCaseFather = t.type({
  children: NestedCaseChildren,
});
`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for nested struct with float64 field", func() {
		type NestedCaseChildren struct {
			Salary float64 `json:"salary"`
		}
		type NestedCaseFather struct {
			Children NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
const NestedCaseChildren = t.type({
  salary: t.number,
});

const NestedCaseFather = t.type({
  children: NestedCaseChildren,
});
`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for nested struct with uint field", func() {
		type NestedCaseChildren struct {
			Count uint `json:"count"`
		}
		type NestedCaseFather struct {
			Children NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
const NestedCaseChildren = t.type({
  count: t.number,
});

const NestedCaseFather = t.type({
  children: NestedCaseChildren,
});
`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for nested struct with pointer field", func() {
		type NestedCaseChildren struct {
			ZipCode *string `json:"zip_code"`
		}
		type NestedCaseFather struct {
			Children NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
const NestedCaseChildren = t.type({
  zip_code: t.union([t.string, t.undefined]),
});

const NestedCaseFather = t.type({
  children: NestedCaseChildren,
});
`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for nested struct with slice field", func() {
		type NestedCaseChildren struct {
			Tags []string `json:"tags"`
		}
		type NestedCaseFather struct {
			Children NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
const NestedCaseChildren = t.type({
  tags: t.array(t.string),
});

const NestedCaseFather = t.type({
  children: NestedCaseChildren,
});
`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for nested struct with array of pointers", func() {
		type NestedCaseChildren struct {
			Scores []*int `json:"scores"`
		}
		type NestedCaseFather struct {
			Children NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
const NestedCaseChildren = t.type({
  scores: t.array(t.union([t.number, t.undefined])),
});

const NestedCaseFather = t.type({
  children: NestedCaseChildren,
});
`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for nested struct with father having array as children", func() {
		type NestedCaseChildren struct {
			Scores []*int `json:"scores"`
		}
		type NestedCaseFather struct {
			Children []NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
const NestedCaseChildren = t.type({
  scores: t.array(t.union([t.number, t.undefined])),
});

const NestedCaseFather = t.type({
  children: t.array(NestedCaseChildren),
});
`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for nested struct with father having pointer array as children", func() {
		type NestedCaseChildren struct {
			Scores []*int `json:"scores"`
		}
		type NestedCaseFather struct {
			Children []*NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
const NestedCaseChildren = t.type({
  scores: t.array(t.union([t.number, t.undefined])),
});

const NestedCaseFather = t.type({
  children: t.array(t.union([NestedCaseChildren, t.undefined])),
});
`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for nested struct with father having pointer array as children, optional array enabled", func() {
		type NestedCaseChildren struct {
			Scores []*int `json:"scores"`
		}
		type NestedCaseFather struct {
			Children []*NestedCaseChildren `json:"children"`
		}

		user := NestedCaseFather{}
		generator := usecase.NewIoTsGenerator(usecase.TypeScriptGeneratorOptions{TreatArraysAsOptional: true})
		result, err := generator.Generate(user)

		expected := `
const NestedCaseChildren = t.type({
  scores: t.union([t.array(t.union([t.number, t.undefined])), t.undefined]),
});

const NestedCaseFather = t.type({
  children: t.union([t.array(t.union([NestedCaseChildren, t.undefined])), t.undefined]),
});
`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})
})

var _ = Describe("IO-TS:Complex Cases", func() {
	It("Hard Complete Test: should convert a complete struct with at least 4 children depth, multiple fields, arrays, etc.", func() {
		type Weapon struct {
			Name        string  `json:"name"`
			Damage      int     `json:"damage"`
			Enchanted   bool    `json:"enchanted"`
			Enchantment *string `json:"enchantment"`
		}

		type Inventory struct {
			Weapons   []Weapon `json:"weapons"`
			Gold      int      `json:"gold"`
			Lockpicks *int     `json:"lockpicks"`
			Potions   []string `json:"potions"`
		}

		type QuestStatus struct {
			Active   bool   `json:"active"`
			Progress string `json:"progress"`
		}

		type Quest struct {
			Title  string      `json:"title"`
			Status QuestStatus `json:"status"`
		}

		type Skill struct {
			Name  string `json:"name"`
			Level int    `json:"level"`
		}

		type Character struct {
			Name      string    `json:"name"`
			Race      string    `json:"race"`
			Health    int       `json:"health"`
			Stamina   *int      `json:"stamina"`
			Inventory Inventory `json:"inventory"`
			Quests    []Quest   `json:"quests"`
			Skills    []Skill   `json:"skills"`
		}

		user := Character{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
const Weapon = t.type({
  name: t.string,
  damage: t.number,
  enchanted: t.boolean,
  enchantment: t.union([t.string, t.undefined]),
});

const Inventory = t.type({
  weapons: t.array(Weapon),
  gold: t.number,
  lockpicks: t.union([t.number, t.undefined]),
  potions: t.array(t.string),
});

const QuestStatus = t.type({
  active: t.boolean,
  progress: t.string,
});

const Quest = t.type({
  title: t.string,
  status: QuestStatus,
});

const Skill = t.type({
  name: t.string,
  level: t.number,
});

const Character = t.type({
  name: t.string,
  race: t.string,
  health: t.number,
  stamina: t.union([t.number, t.undefined]),
  inventory: Inventory,
  quests: t.array(Quest),
  skills: t.array(Skill),
});
`

		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

})

var _ = Describe("IO-TS:Special Cases", func() {

	It("should generate correct io-ts type for map[string]interface{} field", func() {
		type SpecialCase struct {
			Data map[string]interface{} `json:"data"`
		}
		user := SpecialCase{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
const SpecialCase = t.type({
  data: t.record(t.string, t.unknown),
});
`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for map[string]interface{} with omitempty tag", func() {
		type SpecialCase struct {
			Data map[string]interface{} `json:"data,omitempty"`
		}
		user := SpecialCase{}
		generator := usecase.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
const SpecialCase = t.type({
  data: t.union([t.record(t.string, t.unknown), t.undefined]),
});
`

		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})
})

// Handle with map[string]interface{} fields
// Bug: When 	NetworkFunctionLifecycleManagement map[string]interface{}            `json:"networkFunctionLifecycleManagement,omitempty" bson:"networkFunctionLifecycleManagement"` it produces
// networkFunctionLifecycleManagement, omitempty: t.unknown,
