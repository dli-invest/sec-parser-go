package utils

import (
	"unicode"
)

func IsStringAlpha(str string) bool {
	for _, c := range str {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}
