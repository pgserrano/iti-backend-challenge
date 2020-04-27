package password

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	IsValid bool
	Errs []string
}

func createResponse (errs []string ) string {

	var response Result
	if errs != nil {
		response = Result{
			IsValid: false,
			Errs: errs,
		}
	}else {
		response = Result{
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