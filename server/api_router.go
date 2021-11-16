package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yeyudekuangxiang/imagedesign/controller/api"
)

func apiRouter(router *gin.Engine) {
	router.Static("/web", "./web")
	router.POST("api/text", format(api.GetText))
	router.POST("api/code", format(api.GetCode))
	apiRouter := router.Group("/api")
	apiRouter.Use(throttle())
	apiRouter.Use(mustAuth())
	{
		apiRouter.GET("/user", format(api.DefaultUserController.GetUserInfo))
	}
}
