package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pgserrano/iti-backend-challenge/configs/enviroment"
	"github.com/pgserrano/iti-backend-challenge/internal/password"
	"net/http"
)

func InitRoutes(){
	router := SetupRouter()
	router.Run(enviroment.GetSystemPort())
}

func SetupRouter() *gin.Engine{

	router := gin.Default()

	users := router.Group("/users")
		password := users.Group("/passwords")
			validations := password.Group("/validations")
				validations.POST("isValid", handlerPasswordValidation)

	return router
}

func handlerPasswordValidation(c *gin.Context){
	buf := make([]byte, 1024)
	body, _ := c.Request.Body.Read(buf)
	passwd := string(buf[0:body])
	c.String(http.StatusOK, password.IsValid(passwd))
}