package utils

import (
	"strings"
	"unicode"
)

func IsUppercase(s string) bool {
	if IsDigit(s) {
		return false
	}

	return s == strings.ToUpper(s)
}

func IsDigit(s string) bool {
	if len(s) == 0 {
		return false // Empty string is not considered a digit
	}

	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}
