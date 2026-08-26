package reflection

import (
	"reflect"
	"testing"
)

const (
	TEST1  = "Test Walk 1 : Input with one string"
	TEST2  = "Test Walk 2 : Input with a two string"
	TEST3  = "Test Walk 3 : Input with not only string member"
	TEST4  = "Test Walk 4 : Input with nested struct member"
	TEST5  = "Test Walk 5 : Input of type pointer structure"
	TEST6  = "Test Walk 6 : Input of type string slice"
	TEST7  = "Test Walk 7 : Input using struct Array"
	TEST8  = "Test Walk 8 : Input using Map"
	TEST9  = "Test Walk 9 : Input using chan string"
	TEST10 = "Test Walk 10 : Input using func()(string, string)"
)

func TestWalk(t *testing.T) {

	testCases := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
		{
			Name:     TEST1,
			Input:    struct{ Name string }{Name: "Chris"},
			Expected: []string{"Chris"},
		},
		{
			Name: TEST2,
			Input: struct {
				FirstName  string
				SecondName string
			}{"Chris", "Walter"},
			Expected: []string{"Chris", "Walter"},
		},
		{
			Name: TEST3,
			Input: struct {
				Age  int
				Name string
			}{15, "Chris"},
			Expected: []string{"Chris"},
		},
		{
			Name: TEST4,
			Input: struct {
				Name   string
				Adress struct {
					Number int
					Street string
				}
			}{"Chris", struct {
				Number int
				Street string
			}{12, "Triomphe Street"}},
			Expected: []string{"Chris", "Triomphe Street"},
		},
		{
			Name: TEST5,
			Input: &struct {
				Name string
			}{Name: "Chris"},
			Expected: []string{"Chris"},
		},
		{
			Name:     TEST6,
			Input:    []string{"Chris", "Walter"},
			Expected: []string{"Chris", "Walter"},
		},
		{
			Name: TEST7,
			Input: [3]struct {
				Name string
				Date int
			}{{"Chris", 12}, {"Walter", 34}, {"Leebron", 56}},
			Expected: []string{"Chris", "Walter", "Leebron"},
		},
		{
			Name: TEST8,
			Input: map[string]string{
				"Prenom1": "Walter",
				"Prenom2": "Chris",
			},
			Expected: []string{"Walter", "Chris"},
		},
		{
			Name:     TEST9,
			Input:    make(chan string),
			Expected: []string{"Walter", "Chris"},
		},
		{
			Name:     TEST10,
			Input:    func() (string, string) { return "Walter", "Chris" },
			Expected: []string{"Walter", "Chris"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got := []string{}

			if testCase.Name == TEST9 {
				go func() {
					testCase.Input.(chan string) <- "Walter"
					testCase.Input.(chan string) <- "Chris"

					close(testCase.Input.(chan string))
				}()
			}

			walk(testCase.Input, func(input string) {
				got = append(got, input)
			})

			assertLen(t, len(got), len(testCase.Expected))

			switch testCase.Name {
			case TEST8:
				assertSliceContainSameAllElements(t, got, testCase.Expected)
			default:
				if !reflect.DeepEqual(got, testCase.Expected) {
					t.Errorf("Got : %v != want %v", got, testCase.Expected)
				}
			}
		})
	}
}

func assertSliceContainSameAllElements(t *testing.T, input, expected []string) {
	stringPresent := struct {
		Element string
		Present bool
	}{}

	for _, expectedElement := range expected {
		stringPresent.Element = expectedElement
		for _, gotElement := range input {
			if gotElement == expectedElement {
				stringPresent.Present = true
				break
			}
		}
		if !stringPresent.Present {
			t.Errorf("Got : %v != want %v, expected element not present %s", input, expected, stringPresent.Element)
			break
		}
	}
}

func assertLen(t *testing.T, inputLength, expectedLength int) {
	if inputLength != expectedLength {
		t.Errorf("Length are not equal, input Length : %d, expexted length : %d ", inputLength, expectedLength)
	}
}
