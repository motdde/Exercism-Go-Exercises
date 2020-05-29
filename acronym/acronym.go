// Package acronym generate an Acronyn from a string
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate returns the Acronym of words in a sentense.
func Abbreviate(s string) string {

	re := regexp.MustCompile("[A-Z]+[a-z']*|[a-z]+")
	words := re.FindAllString(s, -1)

	var abbr string

	for _, word := range words {
		abbr += strings.ToUpper(string(word[0]))
	}
	return abbr
}
