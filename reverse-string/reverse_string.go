// Package reverse provide method to reverse a UTF strings
package reverse

// func Reverse(s string) string {
// 	runes := []rune(s)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 	  runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
//   }

//Reverse returns reversed string
func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
