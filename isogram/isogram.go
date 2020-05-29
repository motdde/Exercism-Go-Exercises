//Package isogram check if a word is an isogram
package isogram

import "unicode"

//IsIsogram returns a boolean
func IsIsogram(words string) bool {
	checker := make(map[rune]bool)
	for _, v := range words {
		v = unicode.ToLower(v)
		if v != ' ' && v != '-' {
			if checker[v] {
				return false
			}
			checker[v] = true
		}
	}
	return true
}
