package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pgserrano/iti-backend-challenge/configs/enviroment"
	"github.com/pgserrano/iti-backend-challenge/internal/password"
	"github.com/pgserrano/iti-backend-challenge/internal/request"
	"net/http"
	"strconv"
)

func InitRoutes(){
	router := SetupRouter()
	router.Run(enviroment.GetSystemPort())
}

func SetupRouter() *gin.Engine{

	router := gin.Default()

	users := router.Group("/users")
		password := users.Group("/passwords")
			password.POST("/validate", handlerPasswordValidation)


	return router
}

func handlerPasswordValidation(c *gin.Context){
	body := request.ParserBody(c.Request.Body)
	isValid := password.IsValid(body)
	c.String(http.StatusOK, strconv.FormatBool(isValid))
}