//Package accumulate contains a method to map function over a string
package accumulate

//Accumulate return a slice of string
func Accumulate(arr []string, apply func(string) string) []string {
	res := make([]string, len(arr))
	for index, item := range arr {
		res[index] = apply(item)
	}
	return res
}
