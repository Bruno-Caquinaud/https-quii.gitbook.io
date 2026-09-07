package propertybasedtest

import "strings"

type RomanNumeral struct {
	numeral uint16
	roman   string
}

var RomanNumeralTable = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func NumeralRomanConverter(number uint16) string {
	var romanNumber strings.Builder

	for _, romanNumeralelement := range RomanNumeralTable {
		for number >= romanNumeralelement.numeral {
			romanNumber.WriteString(romanNumeralelement.roman)
			number -= romanNumeralelement.numeral
		}
	}

	return romanNumber.String()
}

func RomanNumeralConverter(number string) uint16 {
	var numeralNumber uint16
	for _, romanNumeralelement := range RomanNumeralTable {
		for len(number) > 0 && len(romanNumeralelement.roman) <= len(number) && strings.HasPrefix(number, romanNumeralelement.roman) {
			numeralNumber += romanNumeralelement.numeral
			number = strings.TrimPrefix(number, romanNumeralelement.roman)
		}
	}

	return numeralNumber
}
