package atbash

import "strings"

func Atbash(s string) string {
	s = strings.ToLower(s)
	newText := ""
	count := 0
	for i := range s {
		if s[i] <= 'z' && s[i] >= 'a' {
			byteSisa := 'z' - s[i]
			newText += string('a' + byteSisa)
			count++
		} else if s[i] >= '0' && s[i] <= '9'{
			newText += string(s[i])
			count++
		} else {
			continue
		}

		if count == 5 {
			count = 0
			newText += " "
		}
		}
		return strings.TrimSuffix(newText, " ")
}
