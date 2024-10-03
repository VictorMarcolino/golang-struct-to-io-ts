# Golang Struct to io-ts Generator

This project provides a generator that automatically converts Go structs into `io-ts` types. This tool is particularly useful for projects that need to ensure type safety between backend Go services and frontend TypeScript applications by generating shared types based on Go struct definitions.

## Features

- **io-ts Generator**: Converts Go structs into `io-ts` runtime types, ensuring type-safe data validation in JavaScript/TypeScript.
- **Optional Fields Handling**: Supports pointer and array types, marking fields as optional in `io-ts` when appropriate.
- **Nested Structs**: Recursively generates types for deeply nested Go structs.
- **Inlined Fields**: Supports Go struct fields that are inlined using the `json:",inline"` tag.
- **Special Cases Handling**: Handles special cases like `map[string]interface{}` by generating appropriate `io-ts` types.

## Project Structure

```
├── utils/
│   └── utils.go                 # Utility functions
└── generators/
    ├── generate-io-ts.go        # io-ts type generator
    ├── generate-io-ts_test.go   # io-ts type generator tests
    └── usecase_test.go          # Test runner configuration
```

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/VictorMarcolino/golang-struct-to-io-ts.git
   ```
2. Navigate to the project directory:
   ```bash
   cd golang-struct-to-io-ts
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

### Generate `io-ts` Types

You can generate `io-ts` types by using `IoTsGenerator`.

```go
package main

import (
    "fmt"
    "github.com/VictorMarcolino/golang-struct-to-io-ts/generators"
)

func main() {
    type User struct {
        Name  string `json:"name"`
        Age   int    `json:"age"`
        Email string `json:"email"`
    }

    generator := generators.NewIoTsGenerator()
    result, err := generator.Generate(User{})

    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println(result)
    }
}
```

This will output:

```typescript
import * as t from 'io-ts';

export const UserC = t.type({
  name: t.string,
  age: t.number,
  email: t.string,
});
export type User = t.TypeOf<typeof UserC>;
```

### Handling Optional Fields

Fields that are pointers in Go will be marked as optional in `io-ts`. Additionally, you can pass the `TreatArraysAsOptional` option to the generator to mark arrays as optional if needed.

```go
generator := generators.NewIoTsGenerator(generators.TypeScriptGeneratorOptions{TreatArraysAsOptional: true})
```

### Handling Inlined Fields

The generator supports Go struct fields that are inlined using the `json:",inline"` tag. Inlined fields will have their fields merged into the parent struct in the generated `io-ts` type.

## Running Tests

The project uses [Ginkgo](https://onsi.github.io/ginkgo/) and [Gomega](https://onsi.github.io/gomega/) for testing. You can run the tests by executing:

```bash
ginkgo run -r
```

The test suite covers a variety of cases, including simple types, nested structs, pointer fields, inlined fields, and more complex Go structs with slices and maps.

## Example Tests

### Simple Cases

Tests for basic Go types like `int`, `string`, and `bool`.

```go
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
```

### Nested Structs

Tests for generating `io-ts` types for nested Go structs.

```go
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
```

### Complex Test

A complex test with deeply nested structs, optional fields, arrays, and pointers.

```go
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
    Expect(err).To(BeNil())
    Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
})
```

### Handling Inlined Fields

Tests for structs with inlined fields using the `json:",inline"` tag.

```go
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
    Expect(err).To(BeNil())
    Expect(utils.NormalizeWhitespace(result)).To(Equal(utils.NormalizeWhitespace(expected)))
})
```

## Contributing

Feel free to submit issues or pull requests to improve the functionality, add more features, or fix bugs.

## License

This project is licensed under the MIT License.