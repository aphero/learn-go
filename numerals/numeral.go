package numerals

import "strings"

func ConvertToArabic(roman string) (arabic uint16) {
	arabic = 0
	for _, num := range allRomanNumerals {
		for strings.HasPrefix(roman, num.Symbol) {
			arabic += num.Value
			roman = strings.TrimPrefix(roman, num.Symbol)
		}
	}
	return
}

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic uint16) string {
	// strings.Builder is more optimized for memory performance than just manually editing a string
	var roman strings.Builder

	for _, num := range allRomanNumerals {
		for arabic >= num.Value {
			roman.WriteString(num.Symbol)
			arabic -= num.Value
		}
	}

	return roman.String()
}
