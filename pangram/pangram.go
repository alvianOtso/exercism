package pangram

import "strings"

func IsPangram(input string) bool {
	lowerInput := strings.ToLower(input)

	for char := 'a'; char <= 'z'; char++ {
		if !strings.ContainsRune(lowerInput, char){
			return false
		}
	}

	return true
}
