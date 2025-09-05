package fixtures

type ExampleString string

const (
	ExampleString1   ExampleString = "1"
	ExampleStringTwo ExampleString = "2"
	ExampleString3   ExampleString = "3"
)

type ExampleInt int

const (
	Code1   ExampleInt = 1
	CodeTwo ExampleInt = 2
)

type Example struct {
	ExampleString      ExampleString   `json:"exampleString"`
	ExampleInt         ExampleInt      `json:"exampleInt"`
	ExampleIntArray    []ExampleInt    `json:"exampleIntArray"`
	ExampleStringArray []ExampleString `json:"exampleStringArray"`
}

type RecursionExample struct {
	RecursionExample                *RecursionExample   `json:"recursionExample"`
	RecursionExampleArrayOfPointers []*RecursionExample `json:"recursionExampleArrayOfPointers"`
	RecursionExampleArray           []RecursionExample  `json:"recursionExampleArray"`
	ExampleString                   string              `json:"exampleString"`
	ExampleInt                      int                 `json:"exampleInt"`
}

type AtExample struct {
	AtExample string `json:"@atExample"`
}
