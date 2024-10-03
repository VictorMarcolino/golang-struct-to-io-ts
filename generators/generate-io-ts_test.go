package generators_test

import (
	"github.com/VictorMarcolino/golang-struct-to-io-ts/generators"
	"github.com/VictorMarcolino/golang-struct-to-io-ts/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("IO-TS:Simple Cases", func() {
	It("should generate correct io-ts type for int field", func() {
		type SimpleCase struct {
			Age int `json:"age"`
		}
		user := SimpleCase{}
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `import * as t from 'io-ts';

export const SimpleCaseC = t.type({ age: t.number, });
export type SimpleCase = t.TypeOf<typeof SimpleCaseC>;`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for string field", func() {
		type SimpleCase struct {
			Name string `json:"name"`
		}
		user := SimpleCase{}
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `import * as t from 'io-ts';

export const SimpleCaseC = t.type({ name: t.string, });
export type SimpleCase = t.TypeOf<typeof SimpleCaseC>;`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for bool field", func() {
		type SimpleCase struct {
			IsActive bool `json:"is_active"`
		}
		user := SimpleCase{}
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `import * as t from 'io-ts';

export const SimpleCaseC = t.type({ is_active: t.boolean, });
export type SimpleCase = t.TypeOf<typeof SimpleCaseC>;`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for float64 field", func() {
		type SimpleCase struct {
			Price float64 `json:"price"`
		}
		user := SimpleCase{}
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `import * as t from 'io-ts';

export const SimpleCaseC = t.type({ price: t.number, });
export type SimpleCase = t.TypeOf<typeof SimpleCaseC>;`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for uint field", func() {
		type SimpleCase struct {
			Count uint `json:"count"`
		}
		user := SimpleCase{}
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `import * as t from 'io-ts';

export const SimpleCaseC = t.type({ count: t.number, });
export type SimpleCase = t.TypeOf<typeof SimpleCaseC>;`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for pointer string field", func() {
		type SimpleCase struct {
			ZipCode *string `json:"zip_code"`
		}
		user := SimpleCase{}
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `import * as t from 'io-ts';

export const SimpleCaseC = t.type({ zip_code: t.union([t.string, t.undefined]), });
export type SimpleCase = t.TypeOf<typeof SimpleCaseC>;`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for pointer int field", func() {
		type SimpleCase struct {
			Score *int `json:"score"`
		}
		user := SimpleCase{}
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `import * as t from 'io-ts';

export const SimpleCaseC = t.type({ score: t.union([t.number, t.undefined]), });
export type SimpleCase = t.TypeOf<typeof SimpleCaseC>;`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for slice field", func() {
		type SimpleCase struct {
			Tags []string `json:"tags"`
		}
		user := SimpleCase{}
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `import * as t from 'io-ts';

export const SimpleCaseC = t.type({ tags: t.array(t.string), });
export type SimpleCase = t.TypeOf<typeof SimpleCaseC>;`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for slice of pointers", func() {
		type SimpleCase struct {
			Scores []*int `json:"scores"`
		}
		user := SimpleCase{}
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)
		expected := `import * as t from 'io-ts';

export const SimpleCaseC = t.type({ scores: t.array(t.union([t.number, t.undefined])), });
export type SimpleCase = t.TypeOf<typeof SimpleCaseC>;`
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
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const NestedCaseChildrenC = t.type({
  age: t.number,
});
export type NestedCaseChildren = t.TypeOf<typeof NestedCaseChildrenC>;

export const NestedCaseFatherC = t.type({
  children: NestedCaseChildrenC,
});
export type NestedCaseFather = t.TypeOf<typeof NestedCaseFatherC>;
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
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const NestedCaseChildrenC = t.type({
  name: t.string,
});
export type NestedCaseChildren = t.TypeOf<typeof NestedCaseChildrenC>;

export const NestedCaseFatherC = t.type({
  children: NestedCaseChildrenC,
});
export type NestedCaseFather = t.TypeOf<typeof NestedCaseFatherC>;
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
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const NestedCaseChildrenC = t.type({
  is_active: t.boolean,
});
export type NestedCaseChildren = t.TypeOf<typeof NestedCaseChildrenC>;

export const NestedCaseFatherC = t.type({
  children: NestedCaseChildrenC,
});
export type NestedCaseFather = t.TypeOf<typeof NestedCaseFatherC>;
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
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const NestedCaseChildrenC = t.type({
  salary: t.number,
});
export type NestedCaseChildren = t.TypeOf<typeof NestedCaseChildrenC>;

export const NestedCaseFatherC = t.type({
  children: NestedCaseChildrenC,
});
export type NestedCaseFather = t.TypeOf<typeof NestedCaseFatherC>;
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
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const NestedCaseChildrenC = t.type({
  count: t.number,
});
export type NestedCaseChildren = t.TypeOf<typeof NestedCaseChildrenC>;

export const NestedCaseFatherC = t.type({
  children: NestedCaseChildrenC,
});
export type NestedCaseFather = t.TypeOf<typeof NestedCaseFatherC>;
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
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const NestedCaseChildrenC = t.type({
  zip_code: t.union([t.string, t.undefined]),
});
export type NestedCaseChildren = t.TypeOf<typeof NestedCaseChildrenC>;

export const NestedCaseFatherC = t.type({
  children: NestedCaseChildrenC,
});
export type NestedCaseFather = t.TypeOf<typeof NestedCaseFatherC>;
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
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const NestedCaseChildrenC = t.type({
  tags: t.array(t.string),
});
export type NestedCaseChildren = t.TypeOf<typeof NestedCaseChildrenC>;

export const NestedCaseFatherC = t.type({
  children: NestedCaseChildrenC,
});
export type NestedCaseFather = t.TypeOf<typeof NestedCaseFatherC>;
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
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const NestedCaseChildrenC = t.type({
  scores: t.array(t.union([t.number, t.undefined])),
});
export type NestedCaseChildren = t.TypeOf<typeof NestedCaseChildrenC>;

export const NestedCaseFatherC = t.type({
  children: NestedCaseChildrenC,
});
export type NestedCaseFather = t.TypeOf<typeof NestedCaseFatherC>;
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
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const NestedCaseChildrenC = t.type({
  scores: t.array(t.union([t.number, t.undefined])),
});
export type NestedCaseChildren = t.TypeOf<typeof NestedCaseChildrenC>;

export const NestedCaseFatherC = t.type({
  children: t.array(NestedCaseChildrenC),
});
export type NestedCaseFather = t.TypeOf<typeof NestedCaseFatherC>;
`
		GinkgoWriter.Println("Detailed expected:\n", expected, "\n________________________\nDetailed result:\n", result)
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
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const NestedCaseChildrenC = t.type({
  scores: t.array(t.union([t.number, t.undefined])),
});
export type NestedCaseChildren = t.TypeOf<typeof NestedCaseChildrenC>;

export const NestedCaseFatherC = t.type({
  children: t.array(t.union([NestedCaseChildrenC, t.undefined])),
});
export type NestedCaseFather = t.TypeOf<typeof NestedCaseFatherC>;
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
		generator := generators.NewIoTsGenerator(generators.TypeScriptGeneratorOptions{
			TreatArraysAsOptional: true,
		})
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const NestedCaseChildrenC = t.type({
  scores: t.union([t.array(t.union([t.number, t.undefined])), t.undefined]),
});
export type NestedCaseChildren = t.TypeOf<typeof NestedCaseChildrenC>;

export const NestedCaseFatherC = t.type({
  children: t.union([t.array(t.union([NestedCaseChildrenC, t.undefined])), t.undefined]),
});
export type NestedCaseFather = t.TypeOf<typeof NestedCaseFatherC>;


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
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const WeaponC = t.type({
  name: t.string,
  damage: t.number,
  enchanted: t.boolean,
  enchantment: t.union([t.string, t.undefined]),
});
export type Weapon = t.TypeOf<typeof WeaponC>;

export const InventoryC = t.type({
  weapons: t.array(WeaponC),
  gold: t.number,
  lockpicks: t.union([t.number, t.undefined]),
  potions: t.array(t.string),
});
export type Inventory = t.TypeOf<typeof InventoryC>;

export const QuestStatusC = t.type({
  active: t.boolean,
  progress: t.string,
});
export type QuestStatus = t.TypeOf<typeof QuestStatusC>;

export const QuestC = t.type({
  title: t.string,
  status: QuestStatusC,
});
export type Quest = t.TypeOf<typeof QuestC>;

export const SkillC = t.type({
  name: t.string,
  level: t.number,
});
export type Skill = t.TypeOf<typeof SkillC>;

export const CharacterC = t.type({
  name: t.string,
  race: t.string,
  health: t.number,
  stamina: t.union([t.number, t.undefined]),
  inventory: InventoryC,
  quests: t.array(QuestC),
  skills: t.array(SkillC),
});
export type Character = t.TypeOf<typeof CharacterC>;
`
		GinkgoWriter.Println("Detailed expected:\n", expected, "\n________________________\nDetailed result:\n", result)
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
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const SpecialCaseC = t.type({
  data: t.record(t.string, t.unknown),
});
export type SpecialCase = t.TypeOf<typeof SpecialCaseC>;
`
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for map[string]interface{} with omitempty tag", func() {
		type SpecialCase struct {
			Data map[string]interface{} `json:"data,omitempty"`
		}
		user := SpecialCase{}
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const SpecialCaseC = t.type({
  data: t.union([t.record(t.string, t.unknown), t.undefined]),
});
export type SpecialCase = t.TypeOf<typeof SpecialCaseC>;
`

		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})
})

var _ = Describe("IO-TS:Special Inline Nested Cases", func() {

	It("should generate correct io-ts type for a struct with inline nested structs and optional fields", func() {
		type Attributes struct {
			Strength     int  `json:"strength"`
			Agility      *int `json:"agility,omitempty"` // Optional
			Intelligence int  `json:"intelligence"`
		}

		type Inventory struct {
			Gold    int        `json:"gold"`
			Items   []string   `json:"items"`
			Weapons []struct { // Inline weapons struct
				Name      string `json:"name"`
				Damage    int    `json:"damage"`
				Enchanted bool   `json:"enchanted"`
			} `json:"weapons"`
		}

		type Character struct {
			Name       string     `json:"name"`
			Attributes Attributes `json:",inline"` // This struct is inlined
			Inventory  Inventory  `json:",inline"` // This struct is inlined
		}

		user := Character{}
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const CharacterC = t.type({
  name: t.string,
  strength: t.number,
  agility: t.union([t.number, t.undefined]),
  intelligence: t.number,
  gold: t.number,
  items: t.array(t.string),
  weapons: t.array(t.type({
    name: t.string,
    damage: t.number,
    enchanted: t.boolean,
  })),
});
export type Character = t.TypeOf<typeof CharacterC>;
`
		GinkgoWriter.Println("Detailed expected:\n", expected, "\n________________________\nDetailed result:\n", result)
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for deeply nested inline structs", func() {
		type Attributes struct {
			Strength int  `json:"strength"`
			Agility  *int `json:"agility,omitempty"`
		}

		type Weapon struct {
			Name   string `json:"name"`
			Damage int    `json:"damage"`
		}

		type Inventory struct {
			Gold    int      `json:"gold"`
			Weapons []Weapon `json:"weapons"` // Non-inline nested struct
		}

		type Equipment struct {
			Armor  string `json:"armor"`
			Shield string `json:"shield"`
		}

		type Character struct {
			Name       string     `json:"name"`
			Attributes Attributes `json:",inline"` // Inlined
			Inventory  Inventory  `json:",inline"` // Inlined
			Equipment  Equipment  `json:",inline"` // Inlined nested structs
		}

		user := Character{}
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const WeaponC = t.type({
	name: t.string,
	damage: t.number,
});
export type Weapon = t.TypeOf<typeof WeaponC>;

export const CharacterC = t.type({
	name: t.string,
	strength: t.number,
	agility: t.union([t.number, t.undefined]),
	gold: t.number,
	weapons: t.array(WeaponC),
	armor: t.string,
	shield: t.string,
});
export type Character = t.TypeOf<typeof CharacterC>;
`
		GinkgoWriter.Println("Detailed expected:\n", expected, "\n________________________\nDetailed result:\n", result)

		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for nested structs with array of inlined structs", func() {
		type Quest struct {
			Title       string `json:"title"`
			IsCompleted bool   `json:"is_completed"`
		}

		type Character struct {
			Name   string     `json:"name"`
			Quests []struct { // Inlined quests array with struct
				Quest
				Reward string `json:"reward"`
			} `json:"quests"`
		}

		user := Character{}
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const CharacterC = t.type({
  name: t.string,
  quests: t.array(t.type({
    title: t.string,
    is_completed: t.boolean,
    reward: t.string,
  })),
});
export type Character = t.TypeOf<typeof CharacterC>;
`
		GinkgoWriter.Println("Detailed expected:\n", expected, "\n________________________\nDetailed result:\n", result)
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})

	It("should generate correct io-ts type for struct with multiple inline levels", func() {
		type Location struct {
			City  string `json:"city"`
			State string `json:"state"`
		}

		type Residence struct {
			Home struct {
				Address  string   `json:"address"`
				Location Location `json:",inline"` // Inlined location inside residence
			} `json:",inline"` // Inlined home inside residence
		}

		type Person struct {
			Name      string    `json:"name"`
			Residence Residence `json:",inline"` // Inlined residence inside person
		}

		user := Person{}
		generator := generators.NewIoTsGenerator()
		result, err := generator.Generate(user)

		expected := `
import * as t from 'io-ts';

export const PersonC = t.type({
  name: t.string,
  address: t.string,
  city: t.string,
  state: t.string,
});
export type Person = t.TypeOf<typeof PersonC>;
`
		GinkgoWriter.Println("Detailed expected:\n", expected, "\n________________________\nDetailed result:\n", result)
		Expect(err).To(BeNil())
		Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
	})
})
