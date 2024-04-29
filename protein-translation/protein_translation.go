package protein

import "errors"

var ErrStop error = errors.New("err stop")
var ErrInvalidBase error = errors.New("invalid base")

func FromRNA(rna string) ([]string, error) {
	proteins := []string{}

	for i := 0; i < len(rna); i += 3 {
		codon := rna[i : i+3]
		protein, err := FromCodon(codon)
		if err == ErrInvalidBase {
			return []string{}, err
		}

		if err == ErrStop {
			break
		}

		proteins = append(proteins, protein)
	}

	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	switch codon{
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG": 
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
