package http

import (
	"errors"
	"fmt"
	"io"
)

func ParserBody(body io.Reader) (string, error){

	buf := make([]byte, 1024)
	content, err := body.Read(buf)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("Error on Parser Body")
	}
	return string(buf[0:content]), nil

}