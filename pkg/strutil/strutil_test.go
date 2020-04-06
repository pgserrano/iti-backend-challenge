package strutil

import (
	"fmt"
	"testing"
)


// Esse teste analisa uma coleção de strings válidas e outra inválidas
// Para validar se a String possui ou não determinado tamanho
func TestCheckLength(t *testing.T) {

	baseLength := 9

	okCases := []string{"AbTp9!foo" , "AAAbbbCcC", "xyzabcdef", "xyzabcdefgh" }
	for _, v := range okCases {
		if !(CheckLength(v, baseLength)){
			t.Errorf("Error on CheckLength Test using OK CASES collection. Case tested: %s ", v )
		}
	}

	nokCases := []string{"A", "1", "A1b", "1A2b3C4D", "@" }
	for _, v := range nokCases {
		if CheckLength(v, baseLength){
			t.Errorf("Error on CheckLength Test using NOK CASES collection. Case tested: %s ", v )
		}
	}

	fmt.Println("CheckLength Test OK")
}


// Esse teste analisa uma coleção de strings válidas e outra inválidas
// Para validar se a String possui ou não digitos
func TestContainsDigit(t *testing.T) {

	okCases := []string{"AbTp9!foo", "A1", "1ABC4", "1" }
	for _, v := range okCases {
		if !(ContainsDigit(v)){
			t.Errorf("Error on ContainsDigit Test using OK CASES collection. Case tested: %s ", v )
		}
	}

	nokCases := []string{"AbTpa!foo", "Aa", "ABC", "x", "@#A", "@" }
	for _, v := range nokCases {
		if ContainsDigit(v){
			t.Errorf("Error on ContainsDigit Test using NOK CASES collection. Case tested: %s ", v )
		}
	}

	fmt.Println("ContainsDigit Test OK")
}

// Esse teste analisa uma coleção de strings válidas e outra inválidas
// Para validar se a String possui ou não algum caractere maisuculo
func  TestContainsUpperCase(t *testing.T) {

	okCases := []string{"AbTp9!foo", "A1", "1ABC4", "A" }
	for _, v := range okCases {
		if !(ContainsUpperCase(v)){
			t.Errorf("Error on ContainsUpperCase Test using OK CASES collection. Case tested: %s ", v )
		}
	}

	nokCases := []string{"abtpa!foo", "aa", "abc", "x", "@" }
	for _, v := range nokCases {
		if ContainsUpperCase(v){
			t.Errorf("Error on ContainsUpperCase Test using NOK CASES collection. Case tested: %s ", v )
		}
	}

	fmt.Println("ContainsUpperCase Test OK")
}

// Esse teste analisa uma coleção de strings válidas e outra inválidas
// Para validar se a String possui ou não algum caractere minusculo
func TestContainsLowerCase(t *testing.T) {

	okCases := []string{"AbTp9!foo", "A1a", "1abC4", "a1@" }
	for _, v := range okCases {
		if !(ContainsLowerCase(v)){
			t.Errorf("Error on ContainsLowerCase Test using OK CASES collection. Case tested: %s ", v )
		}
	}

	nokCases := []string{"ABTPA!FOO", "AA", "ABC", "2", "!@" }
	for _, v := range nokCases {
		if ContainsLowerCase(v){
			t.Errorf("Error on ContainsLowerCase Test using NOK CASES collection. Case tested: %s ", v )
		}
	}

	fmt.Println("ContainsLowerCase Test OK")
}

// Esse teste analisa uma coleção de strings válidas e outra inválidas
// Para validar se a String possui ou não algum caractere especial
func TestContainsSpecialCharacter(t *testing.T) {

	okCases := []string{"AbTp9!foo", "A!", "1ABC4@#", "@" }
	for _, v := range okCases {
		if !(ContainsSpecialCharacter(v)){
			t.Errorf("Error on ContainsSpecialCharacter Test using OK CASES collection. Case tested: %s ", v )
		}
	}

	nokCases := []string{"AbTpaAfoo", "Aa", "ABC", "x" }
	for _, v := range nokCases {
		if ContainsSpecialCharacter(v){
			t.Errorf("Error on ContainsSpecialCharacter Test using NOK CASES collection. Case tested: %s ", v )
		}
	}

	fmt.Println("ContainsSpecialCharacter Test OK")
}