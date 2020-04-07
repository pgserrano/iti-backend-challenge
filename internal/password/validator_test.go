package password

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)


// Esse teste analisa uma coleção de strings válidas e outra inválidas
// Para validar se a String é válida ou não como senha
func TestIsValid(t *testing.T) {

	okResponse := "{\"IsValid\":true,\"Errs\":null}"
	okCases := []string{`{"password":"AbTp9!foo"}`, `{"password":"XyzkJ#123"}`, `{"password":"178Mn@fcb"}` , `{"password":"SwROTS3!%&*"}`  }
	for _, v := range okCases {
		assert.Equal(t, okResponse, IsValid(v))
	}

	nokCases := []string{`{"password":"A"}`, `{"password":"b"}`, `{"password":"Ab1#"}` , `{"password":"@"}`}
	for _, v := range nokCases {
		assert.NotEqual(t, okResponse, IsValid(v))
	}

	fmt.Println("IsValid Test OK")
}
