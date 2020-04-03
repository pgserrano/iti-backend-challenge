package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pgserrano/iti-backend-challenge/configs/enviroment"
	"github.com/pgserrano/iti-backend-challenge/internal/password"
	"strconv"
)

func InitRoutes(){

	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	password := route.Group("/passwords")
		validations := password.Group("/validations")
			validations.POST("handlePasswordValidation", handlePasswordValidation)



	route.Run(enviroment.GetSystemPort())

}

func handlePasswordValidation(c *gin.Context){
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	passwd := string(buf[0:num])
	isValid := password.IsValid(passwd)
	c.String(200, strconv.FormatBool(isValid))
}