package main

import (
	"github.com/MuhammadSuryono/module-golang-server/server"
	"github.com/gin-gonic/gin"
	"net/http/pprof"
	"upoader-golang/routes"
)

func main() {
	router := server.ConfigServer()
	api := router.Group("api/v1")
	{
		routes.Route(api)
		api.GET("/debug/pprof", func(context *gin.Context) {
			pprof.Index(context.Writer, context.Request)
		})
	}

	server.RunServer("8082")
}
