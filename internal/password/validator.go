package password

import (
	"encoding/json"
	"github.com/pgserrano/iti-backend-challenge/pkg/strutil"
	"strconv"
)

type PasswordRequest struct {
	password string
}

//TODO Criar uma chain de validações. Como fazer ?  Elixir/F#-like
func IsValid(password string) string {

	passwd := extractPassword(password)
	var errs []string

	minimunLength := 9
	isLengthOk := strutil.CheckLength(passwd, minimunLength)
	if !isLengthOk{
		err := "Quantidade de caracteres menor do que o limite mínimo de " + strconv.Itoa(minimunLength)
		errs = append(errs, err)
	}

	contaisDigit := strutil.ContainsDigit(passwd)
	if !contaisDigit{
		err := "A senha não possui nenhum digito"
		errs = append(errs, err)
	}

	contaisLowerCase := strutil.ContainsLowerCase(passwd)
	if !contaisLowerCase{
		err := "A senha não possui nenhum caractere minusculo"
		errs = append(errs, err)
	}

	containsUpperCase := strutil.ContainsUpperCase(passwd)
	if !containsUpperCase{
		err := "A senha não possui nenhum caractere maiusculo"
		errs = append(errs, err)
	}

	containsSpecialChar := strutil.ContainsSpecialCharacter(passwd)
	if !containsSpecialChar{
		err := "A senha não possui nenhum caractere especial"
		errs = append(errs, err)
	}

	return createResponse(errs)
}

func extractPassword(password string) string {

	var contentMap map[string]string
	if err := json.Unmarshal([]byte(password), &contentMap); err != nil {
		panic(err)
	}
	strPassword := contentMap["password"]
	return strPassword
}


