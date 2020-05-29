//Package strain provides methods that retuns a new collection
package strain

//Ints type
type Ints []int

//Lists type
type Lists [][]int

//Strings type
type Strings []string

//Keep for  Ints type
func (i Ints) Keep(f func(int) bool) (result Ints) {
	for _, k := range i {
		if f(k) {
			result = append(result, k)
		}
	}
	return
}

//Discard for Ints type
func (i Ints) Discard(f func(int) bool) (result Ints) {
	for _, k := range i {
		if !f(k) {
			result = append(result, k)
		}
	}
	return
}

//Keep for  Lists type
func (l Lists) Keep(f func([]int) bool) (result Lists) {
	for _, k := range l {
		if f(k) {
			result = append(result, k)
		}
	}
	return
}

//Keep for  Lists type
func (s Strings) Keep(f func(string) bool) (result Strings) {
	for _, k := range s {
		if f(k) {
			result = append(result, k)
		}
	}
	return
}
