// Package triangle find the type of traingle
package triangle

import "math"

// Kind is a type of string
type Kind int

const (
	//NaT not a triangle
	NaT = iota // not a triangle
	//Equ equilateral triangle
	Equ
	//Iso isosceles triangle
	Iso
	//Sca scalene triangle
	Sca
)

// KindFromSides returns Kind type
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	var isValid bool

	args := []float64{a, b, c}

	for _, v := range args {
		isValid = math.IsInf(v, 0) || v <= 0 || !(a+b >= c && b+c >= a && c+a >= b)
		if isValid {
			k = NaT
			return k
		}
	}

	if a == b && a == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}
	return k
}