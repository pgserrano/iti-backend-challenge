package password

import (
	"fmt"
	"testing"
)


// Esse teste analisa uma coleção de strings válidas e outra inválidas
// Para validar se a String é válida ou não como senha
func TestIsValid(t *testing.T) {

	okCases := []string{"AbTp9!foo" , "XyzkJ#123", "178Mn@fcb", "SwROTS3!%&*" }
	for _, v := range okCases {
		if !(IsValid(v)){
			t.Errorf("Error on IsValid Test using OK CASES collection. Case tested: %s ", v )
		}
	}

	nokCases := []string{"A", "1", "A1b", "1A2b3C4D", "@", "SwROTSE!%&*", "178MnAfcb", "XyzkJ#AbC" }
	for _, v := range nokCases {
		if IsValid(v){
			t.Errorf("Error on IsValid Test using NOK CASES collection. Case tested: %s ", v )
		}
	}

	fmt.Println("IsValid Test OK")
}
