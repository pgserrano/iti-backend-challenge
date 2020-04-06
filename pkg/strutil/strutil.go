package strutil

import (
	"regexp"
	"unicode"
)



func CheckLength(str string, minLength int) bool {

	isLengthOk := len(str) >= minLength

	return isLengthOk
}

func ContainsDigit(str string) bool{

	match, _ := regexp.MatchString("\\d", str)
	return match

}

func ContainsUpperCase(str string) bool {

	for _, char := range str {
		if unicode.IsLetter(char) && unicode.IsUpper(char) {
			return true
		}
	}

	return false
}

func ContainsLowerCase(str string) bool {

	for _, char := range str {
		if unicode.IsLetter(char) && unicode.IsLower(char) {
			return true
		}
	}

	return false
}

func ContainsSpecialCharacter(str string) bool {

	match, _ := regexp.MatchString("\\W", str)
	return match
}
