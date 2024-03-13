package strand

func ToRNA(dna string) string {
	var rna string
	for _, nuc := range dna{
		switch nuc{
		case 'A':
			rna += "U"
		case 'C':
			rna += "G"
		case 'G':
			rna += "C"
		case 'T':
			rna += "A"
		default:
			rna += ""
		}
	}

	return rna
}
