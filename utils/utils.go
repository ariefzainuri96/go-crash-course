package utils

import (
	"math/rand"
	"strings"
	"time"
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

// Initialize a new random source and generator
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandBool() bool {
	return rnd.Intn(2) == 1
}
