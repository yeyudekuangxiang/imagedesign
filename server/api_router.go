package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yeyudekuangxiang/imagedesign/controller/api"
	"net/http"
)

func apiRouter(router *gin.Engine) {
	router.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusFound, "/web")
	})
	router.Static("/web", "./web")
	router.POST("/api/text", format(api.GetText))
	router.POST("/api/code", format(api.GetCode))
}
