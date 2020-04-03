package password

import (
	"github.com/pgserrano/iti-backend-challenge/internal/strutil"
)

//TODO Criar uma chain de validações. Como fazer ?
func IsValid(password string) bool {

	isLengthOk := strutil.CheckLength(password, 9)
	contaisDigit := strutil.ContaisDigit(password)
	contaisLowerCase := strutil.ContaisLowerCase(password)
	containsUpperCase := strutil.ContainsUpperCase(password)
	containsSpecialChar := strutil.SpecialCharacter(password)

	return isLengthOk && contaisDigit && contaisLowerCase && containsUpperCase && containsSpecialChar

}
