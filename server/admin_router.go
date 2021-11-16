package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yeyudekuangxiang/imagedesign/controller/admin"
)

func adminRouter(router *gin.Engine) {
	adminRouter := router.Group("/admin")
	adminRouter.Use(authAdmin())
	{
		adminRouter.GET("/user", format(admin.DefaultUserController.GetUserInfo))
	}
}
