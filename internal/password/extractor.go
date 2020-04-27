package password

import "encoding/json"

type PasswordRequest struct {
	password string
}

func ExtractPassword(password string) string {

	var contentMap map[string]string
	if err := json.Unmarshal([]byte(password), &contentMap); err != nil {
		panic(err)
	}
	strPassword := contentMap["password"]
	return strPassword
}