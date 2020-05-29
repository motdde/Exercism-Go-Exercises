// Package proverb make proverbs from a slice
package proverb

import "fmt"

// Proverb accepts a slice returns a string
func Proverb(rhyme []string) []string {
	var result []string
	if len(rhyme) == 0 {
		return result
	}

	for v := range rhyme {
		if v == len(rhyme)-1 {
			result = append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
			break
		}
		result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[v], rhyme[v+1]))
	}
	return result
}
