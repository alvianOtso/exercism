package logs

import (
	"fmt"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		charUnicode := fmt.Sprintf("%U", char)
		switch charUnicode {
		case "U+2757":
			return "recommendation"
		case "U+1F50D":
			return "search"
		case "U+2600":
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	logRune := []rune(log)
	for i, charRune := range logRune {
		if charRune == oldRune {
			logRune[i] = newRune
		}
	}
	return string(logRune)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	logLength := utf8.RuneCountInString(log)
	return logLength <= limit
}
