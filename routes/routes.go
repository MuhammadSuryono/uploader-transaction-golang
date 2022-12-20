package routes

import (
	"github.com/gin-gonic/gin"
	"upoader-golang/qr"
)

func Route(router *gin.RouterGroup) {
	api := router.Group("transaction")
	{
		api.POST("/upload/qr", qr.QrValidateRequest, qr.InvokeRequest)
	}
}
