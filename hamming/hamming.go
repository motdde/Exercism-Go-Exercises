// Package hamming checks the distance between two DNA strands
package hamming

import "errors"

// Distance takes two string agrument and returns an Int and Bool
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("cant check the distance")
	}
	var num int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			num++
		}
	}
	return num, nil
}
