package main

import (
	"errors"
	"fmt"
)

func FromRNA(rna string) ([]string, error) {
	codons := []string{}
	if len(rna) % 3 != 0 {
		return codons, errors.New("invalid rna")
	}	
	
	for i := 0; i < len(rna); i += 3 {
		codon := rna[i : i+3]
		codons = append(codons, codon)
	}

	return codons, nil
}
func main() {
	b, err := FromRNA("AUGUUUUCUUAAAU")
	fmt.Println(b, err)
}
