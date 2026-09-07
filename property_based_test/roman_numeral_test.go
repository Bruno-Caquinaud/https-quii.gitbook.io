package propertybasedtest

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

var testcases = []struct {
	description string
	input       uint16
	expected    string
}{
	{
		description: "Test Roman Numeral 1 : When input = 1, expected output = I",
		input:       1,
		expected:    "I",
	},
	{
		description: "Test Roman Numeral 2 : When input = 2, expected output = II",
		input:       2,
		expected:    "II",
	},
	{
		description: "Test Roman Numeral 3 : When input = 4, expected output = IV",
		input:       4,
		expected:    "IV",
	},
	{
		description: "Test Roman Numeral 4 : When input = 8, expected output = VIII",
		input:       8,
		expected:    "VIII",
	},
	{
		description: "Test Roman Numeral 5 : When input = 9, expected output = IX",
		input:       9,
		expected:    "IX",
	},
	{
		description: "Test Roman Numeral 6 : When input = 16, expected output = XVI",
		input:       16,
		expected:    "XVI",
	},
	{
		description: "Test Roman Numeral 7 : When input = 33, expected output = XXXIII",
		input:       33,
		expected:    "XXXIII",
	},
	{
		description: "Test Roman Numeral 8 : When input = 46, expected output = XLVI",
		input:       46,
		expected:    "XLVI",
	},
	{
		description: "Test Roman Numeral 9 : When input = 68, expected output = LXVIII",
		input:       68,
		expected:    "LXVIII",
	},
	{
		description: "Test Roman Numeral 10 : When input = 99, expected output = XCIX",
		input:       99,
		expected:    "XCIX",
	},
	{
		description: "Test Roman Numeral 11 : When input = 135, expected output = CXXXV",
		input:       135,
		expected:    "CXXXV",
	},
	{
		description: "Test Roman Numeral 12 : When input = 456, expected output = CDLVI",
		input:       456,
		expected:    "CDLVI",
	},
	{
		description: "Test Roman Numeral 13 : When input = 852, expected output = DCCCLII",
		input:       852,
		expected:    "DCCCLII",
	},
	{
		description: "Test Roman Numeral 14 : When input = 999, expected output = CMXCIX",
		input:       999,
		expected:    "CMXCIX",
	},
	{
		description: "Test Roman Numeral 15 : When input = 1200, expected output = MCC",
		input:       1200,
		expected:    "MCC",
	}, {
		description: "Test Roman Numeral 16 : When input = 99, expected output = XCIX",
		input:       99,
		expected:    "XCIX",
	},
}

func TestNumeralRomanConverter(t *testing.T) {

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			romanNumber := NumeralRomanConverter(testcase.input)

			if romanNumber != testcase.expected {
				t.Errorf("Roman Number : %s is not equal to expected one :%s", romanNumber, testcase.expected)
			}
		})
	}

}

func TestRomanNumeralConverter(t *testing.T) {

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			romanNumber := RomanNumeralConverter(testcase.expected)

			if romanNumber != testcase.input {
				t.Errorf("Roman Number : %d is not equal to expected one :%d", testcase.input, romanNumber)
			}
		})
	}
}

func TestRomanNumeralPropertyBasedTesting(t *testing.T) {
	t.Run("Property Based Test 1 : Romanconversion of NumeralConversion number == number",
		func(t *testing.T) {
			assertion := func(number uint16) bool {
				return RomanNumeralConverter(NumeralRomanConverter(number)) == number
			}

			config := &quick.Config{
				MaxCount: 1000,
				Values: func(args []reflect.Value, r *rand.Rand) {
					args[0] = reflect.ValueOf(uint16(r.Intn(4000)))
				},
			}

			if err := quick.Check(assertion, config); err != nil {
				t.Errorf("Failed to convert this number : %v", err)
			}
		})
	t.Run("Property Based Test 2 : No more than 3 consecutive roman symbol",
		func(t *testing.T) {
			assertion := func(number uint16) bool {
				roman := NumeralRomanConverter(number)
				if len(roman) > 3 {
					for i := 0; i+3 < len(roman); i = i + 1 {
						if i == i+1 && i == i+2 && i == i+3 {
							return false
						}
					}
				}
				return true
			}

			config := &quick.Config{
				MaxCount: 10000,
				Values: func(args []reflect.Value, r *rand.Rand) {
					args[0] = reflect.ValueOf(uint16(r.Intn(4000)))
				},
			}

			if err := quick.Check(assertion, config); err != nil {
				t.Errorf("Failed to convert this number : %v", err)
			}
		})
}
