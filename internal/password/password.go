package password

import (
	"encoding/json"
	"fmt"
	"github.com/pgserrano/iti-backend-challenge/pkg/strutil"
	"strconv"
)

type PasswordResponse struct {
	IsValid bool
	Errs []string
}

//TODO Criar uma chain de validações. Como fazer ?  Elixir/F#-like
func IsValid(password string) string {

	var errs []string

	minimunLength := 9
	isLengthOk := strutil.CheckLength(password, minimunLength)
	if !isLengthOk{
		err := "Quantidade de caracteres menor do que o limite mínimo de " + strconv.Itoa(minimunLength)
		errs = append(errs, err)
	}

	contaisDigit := strutil.ContainsDigit(password)
	if !contaisDigit{
		err := "A senha não possui nenhum digito"
		errs = append(errs, err)
	}

	contaisLowerCase := strutil.ContainsLowerCase(password)
	if !contaisLowerCase{
		err := "A senha não possui nenhum caractere minusculo"
		errs = append(errs, err)
	}

	containsUpperCase := strutil.ContainsUpperCase(password)
	if !containsUpperCase{
		err := "A senha não possui nenhum caractere maiusculo"
		errs = append(errs, err)
	}

	containsSpecialChar := strutil.ContainsSpecialCharacter(password)
	if !containsSpecialChar{
		err := "A senha não possui nenhum caractere especial"
		errs = append(errs, err)
	}

	return createResponse(errs)
}


func createResponse (errs []string ) string {

	var response PasswordResponse
	if errs != nil {
		response = PasswordResponse{
					IsValid: false,
					Errs: errs,
		}
	}else {
		response = PasswordResponse{
			IsValid: true,
			Errs: nil,
		}
	}

	responseJson , errJson := json.Marshal(response)
	if(errJson != nil){
		fmt.Println(errJson)
	}

	return string(responseJson)

}