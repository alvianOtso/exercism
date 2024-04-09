package romannumerals

import (
	"errors"
	"strings"
)

var romanNumeric = map[int]byte{
	1000: 'M',
	500:  'D',
	100:  'C',
	50:   'L',
	10:   'X',
	5:    'V',
	1:    'I',
}

var numeric = []int{1000, 500, 100, 50, 10, 5, 1}
var replacement = map[string]string{
	"VIV": "IX",
	"LXL" : "XC",
	"DCD" : "CM",
}

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("invalid input")
	}

	romanString := ""
	idx := 0
	count := 0
	for (input > 0) {
		temp := input - numeric[idx]
		if temp < 0 {
			count = 0
			idx++
			continue
		}

		romanString += string(romanNumeric[numeric[idx]])
		input = temp
		count++

		if count > 3 {
			romanString = romanString[:len(romanString)-3] + string(romanNumeric[numeric[idx-1]])
		}
	}

	for key, value := range replacement {
		romanString = strings.ReplaceAll(romanString, key, value)
	}

	return romanString, nil
}
