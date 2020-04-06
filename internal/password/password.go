package password

import (
	"github.com/pgserrano/iti-backend-challenge/pkg/strutil"
)



//TODO Criar uma chain de validações. Como fazer ?  Elixir/F#-like
func IsValid(password string) bool {

	isLengthOk := strutil.CheckLength(password, 9)
	contaisDigit := strutil.ContainsDigit(password)
	contaisLowerCase := strutil.ContainsLowerCase(password)
	containsUpperCase := strutil.ContainsUpperCase(password)
	containsSpecialChar := strutil.ContainsSpecialCharacter(password)

	return isLengthOk && contaisDigit && contaisLowerCase && containsUpperCase && containsSpecialChar

}
