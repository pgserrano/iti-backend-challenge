package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pgserrano/iti-backend-challenge/configs/enviroment"
	"github.com/pgserrano/iti-backend-challenge/internal/password"
)

func StartServer(){
	router := SetupRouter()
	router.Run(enviroment.GetSystemPort())
}

func SetupRouter() *gin.Engine{

	router := gin.Default()
	password.CreateRouters(router)

	return router
}

