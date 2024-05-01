package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	wordMap := Frequency{}
	reg := regexp.MustCompile(`\w+('\w+)?`)
	
	for _, word := range reg.FindAllString(strings.ToLower(phrase), -1) {
		wordMap[word]++
	}
	return wordMap
}
