package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	oneToOneMap := map[string]int{}
	for k,letters := range in{
		for _, letter := range letters {
			oneToOneMap[strings.ToLower(letter)] = k
		}
	}	

	return oneToOneMap
}
