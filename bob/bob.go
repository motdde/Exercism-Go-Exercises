// Package bob should have a package comment that summarizes what it's about.
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	q := strings.HasSuffix(remark, "?")
	switch {
	case q:
		if containsLetter(remark) && remark == strings.ToUpper(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	case containsLetter(remark) && remark == strings.ToUpper(remark):
		return "Whoa, chill out!"
	case remark == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}

func containsLetter(s string) bool {
	var length int
	for _, caracter := range s {
		if !unicode.IsLetter(caracter) {
			length++
		}
	}

	if length == len(s) {
		return false
	}

	return true
}
