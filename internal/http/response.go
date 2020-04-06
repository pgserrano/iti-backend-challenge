package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleBadRequest(context *gin.Context) {

	context.String(http.StatusBadRequest, "Body com formato inválido")
}


func HandleStatusOK(context *gin.Context, response string){

	context.String(http.StatusOK, response)

}