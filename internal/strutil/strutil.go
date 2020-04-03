package strutil

import (
	"strings"
	"unicode"
)

func CheckLength(str string, minLength int) bool {

	isLengthOk := len(str) >= minLength

	return isLengthOk
}

func ContaisDigit(str string) bool{

	containsDigit := strings.ContainsAny(str, "0123456789")
	return containsDigit

}

func ContainsUpperCase(str string) bool {

	for _, char := range str {
		if unicode.IsLetter(char) && unicode.IsUpper(char) {
			return true
		}
	}

	return false
}

func ContaisLowerCase(str string) bool {

	for _, char := range str {
		if unicode.IsLetter(char) && unicode.IsLower(char) {
			return true
		}
	}

	return false
}

func SpecialCharacter(str string) bool {

	for _, char := range str {
		if unicode.IsSymbol(char) {
			return true
		}
	}

	return false
}