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
			validations.POST("isValidPassword", isValidPassword)



	route.Run(enviroment.GetSystemPort())

}

func isValidPassword(c *gin.Context){
	passwd := "abc"
	isValid := password.IsValidPassword(passwd)
	c.String(200, strconv.FormatBool(isValid))
}