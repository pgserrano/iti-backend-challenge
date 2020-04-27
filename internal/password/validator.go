package password

import (
	"encoding/json"
	"github.com/pgserrano/iti-backend-challenge/pkg/strutil"
	"strconv"
)

type PasswordRequest struct {
	password string
}

var errs []string

//TODO Criar uma chain de validações. Como fazer ?  Elixir/F#-like
func IsValid(password string) string {

	passwd := extractPassword(password)

	minimunLength := 9
	checkConstraint(strutil.CheckLength(passwd, minimunLength),"Quantidade de caracteres menor do que o limite mínimo de " + strconv.Itoa(minimunLength))

	checkConstraint(strutil.ContainsDigit(passwd),  "A senha não possui nenhum digito")

	checkConstraint(strutil.ContainsLowerCase(passwd),  "A senha não possui nenhum caractere minusculo")

	checkConstraint(strutil.ContainsUpperCase(passwd),  "A senha não possui nenhum caractere maiusculo")

	checkConstraint(strutil.ContainsSpecialCharacter(passwd),  "A senha não possui nenhum caractere especial")

	response := createResponse(errs)
	errs = nil

	return response
}

func checkConstraint(constraint bool, msg string) {

	if !constraint {
		errs = append(errs, msg)
	}
}

func extractPassword(password string) string {

	var contentMap map[string]string
	if err := json.Unmarshal([]byte(password), &contentMap); err != nil {
		panic(err)
	}
	strPassword := contentMap["password"]
	return strPassword
}

