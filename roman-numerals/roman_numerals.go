//Package romannumerals converts number to roman numerals
package romannumerals

import (
	"fmt"
)

var arabicToRomanNumeral = []struct {
	arabic       int
	romanNumeral string
}{
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

// ToRomanNumeral return a string or err
func ToRomanNumeral(number int) (result string, err error) {
	if number < 1 || number > 3000 {
		err = fmt.Errorf("Invalid input: %d must be between 1 and 3000", number)
	}

	for _, values := range arabicToRomanNumeral {
		for ; number >= values.arabic; number -= values.arabic {
			result += values.romanNumeral
		}
	}
	return
}
