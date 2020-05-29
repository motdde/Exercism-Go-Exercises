//Package dna count the number of nucleotide
package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[byte]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA []byte

// Counts returs type Histogram
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	for _, v := range d {
		switch v {
		case 'A', 'C', 'G', 'T':
			h[v]++
		default:
			return nil, errors.New("invalid DNA")
		}
	}
	return h, nil
}
