//Package protein provides FronRNA method
package protein

import "errors"

//Error types
var (
	ErrStop        = errors.New("ErrStop")
	ErrInvalidBase = errors.New("ErrInvalidBase")
)

//FromRNA returns ploypeptide sequences
func FromRNA(rna string) (sequence []string, err error) {
	runes := []rune(rna)

	for i := 0; i < len(runes); i += 3 {
		protien, err := FromCodon(string([]rune{runes[i], runes[i+1], runes[i+2]}))
		if err != nil {
			if err == ErrStop {
				break
			}
			return sequence, err
		}
		sequence = append(sequence, protien)
	}
	return
}

// FromCodon to protien
func FromCodon(condons string) (protien string, err error) {
	switch condons {
	case "AUG":
		protien = "Methionine"
	case "UUU", "UUC":
		protien = "Phenylalanine"
	case "UUA", "UUG":
		protien = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protien = "Serine"
	case "UAU", "UAC":
		protien = "Tyrosine"
	case "UGU", "UGC":
		protien = "Cysteine"
	case "UGG":
		protien = "Tryptophan"
	case "UAA", "UAG", "UGA":
		protien = ""
		err = ErrStop
	default:
		protien = ""
		err = ErrInvalidBase
	}
	return
}
