package password

import (
	"github.com/gin-gonic/gin"
	"github.com/pgserrano/iti-backend-challenge/internal/http"
)

func CreateRouters(router *gin.Engine) {

	users := router.Group("/users")
		password := users.Group("/passwords")
			password.POST("/validate", handlerPasswordValidation)

}


func handlerPasswordValidation(c *gin.Context){

	body, err := http.ParserBody(c.Request.Body)
	if err != nil {
		http.HandleBadRequest(c)
	} else {
		http.HandleStatusOK(c, IsValid(body))
	}

}