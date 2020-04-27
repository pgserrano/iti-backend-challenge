package password

import (
	"github.com/pgserrano/iti-backend-challenge/pkg/strutil"
	"strconv"
)



var errs []string

//TODO Criar uma chain de validações. Como fazer ?  Elixir/F#-like
func IsValid(password string) string {

	passwd := ExtractPassword(password)

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



