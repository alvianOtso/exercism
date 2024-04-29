package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {
	newSubject := strings.ToLower(subject)
	anagrams := []string{}
	runeMap := make(map[rune]int)

	for _, rune := range newSubject {
		runeMap[rune]++
	}

	for _, canditate := range candidates {
		anagram := true
		newCandidate := strings.ToLower(canditate)
		if len(subject) != len(canditate) || newCandidate == newSubject {
			continue
		}

		for _, runeSubject := range newSubject {
			counts := strings.Count(newCandidate, string(runeSubject))
			if counts != runeMap[runeSubject] {
				anagram = false
				break
			}
		}

		if anagram {
			anagrams = append(anagrams, canditate)
		}
	}
	return anagrams
}
