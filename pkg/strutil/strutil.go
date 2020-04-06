package strutil

import (
	"regexp"
	"unicode"
)


// CheckLength analisa se uma String tem determinado tamanho
func CheckLength(str string, minLength int) bool {

	isLengthOk := len(str) >= minLength

	return isLengthOk
}

// ContainsDigit analisa se uma String possui algum digito
// Utilizando regex para fazer essa análise
func ContainsDigit(str string) bool{

	match, _ := regexp.MatchString("\\d", str)
	return match

}

// ContainsUpperCase analisa se uma String possui algum caractere maiusculo
// Utilizando runes para fazer essa análise
func ContainsUpperCase(str string) bool {

	for _, char := range str {
		if unicode.IsLetter(char) && unicode.IsUpper(char) {
			return true
		}
	}

	return false
}

// ContainsLowerCase analisa se uma String possui algum caractere minusculo
// Utilizando runes para fazer essa análise
func ContainsLowerCase(str string) bool {

	for _, char := range str {
		if unicode.IsLetter(char) && unicode.IsLower(char) {
			return true
		}
	}

	return false
}

// ContainsSpecialCharacter analisa se uma String possui algum caractere especial
// Utilizando regex para fazer essa análise
func ContainsSpecialCharacter(str string) bool {

	match, _ := regexp.MatchString("\\W", str)
	return match
}
