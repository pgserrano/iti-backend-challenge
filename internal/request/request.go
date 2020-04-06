package request

import (
	"fmt"
	"io"
)

func ParserBody(body io.Reader) string{

	buf := make([]byte, 1024)
	content, err := body.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	return string(buf[0:content])

}