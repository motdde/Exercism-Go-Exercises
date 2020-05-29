//Package strand converts DNA to RNA
package strand

var conv = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

//ToRNA retunds a string
func ToRNA(dna string) string {
	rna := ""
	for _,v := range dna {
		rna += conv[v]
	}
	return rna
}

// var rna strings.Builder
// rna.Grow(len(dna))
// for _, v := range dna {
// 	fmt.Fprintf(&rna, "%v", string(conv[v]))
// }
// return rna.String()