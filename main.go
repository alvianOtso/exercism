package main

import (
	"fmt"
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	isogram := true
	for _, r := range word {
		fmt.Println(string(r), strings.Count(word, string(r)))
		if strings.Count(word, string(r)) > 1 {
			isogram = false
			break
		}
	}

	return isogram
}
func main() {
	b := IsIsogram("six-year-old")
	fmt.Println(b)
}
