package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)


func ParserBody(context *gin.Context) (string, error) {

	content, err := ioutil.ReadAll(context.Request.Body)
	isValid := json.Valid(content)
	if err != nil  || !isValid{
		fmt.Println(err)
		return "", errors.New("Error on Parser Body")
	}
	return string(content), nil

}