package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yeyudekuangxiang/imagedesign/internal/util"
)

var DefaultUserController = UserController{}

type UserController struct {
}

func (UserController) GetUserInfo(c *gin.Context) (gin.H, error) {
	user := util.GetAuthUser(c)
	return gin.H{
		"user": user,
	}, nil
}
