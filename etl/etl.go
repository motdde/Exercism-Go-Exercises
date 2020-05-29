//Package etl provides method for transforming legacy point systems to new
package etl

import (
	"strings"
)

//Transform return the new data
func Transform(legacy map[int][]string) (new map[string]int) {
	new = make(map[string]int)
	for point, letters := range legacy {
		for _, letter := range letters {
			new[strings.ToLower(letter)] = point
		}
	}
	return
}
