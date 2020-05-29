//Package raindrops converts integer to rain sounds
package raindrops

import "strconv"

//Convert return a string
func Convert(num int) string {
	var result string

	if num%3 == 0 {
		result += "Pling"
	}
	if num%5 == 0 {
		result += "Plang"
	}
	if num%7 == 0 {
		result += "Plong"
	}
	if len(result) == 0 {
		return strconv.Itoa(num)
	}
	return result
}
