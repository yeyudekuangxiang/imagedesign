package util

import (
	"github.com/gin-gonic/gin"
	"github.com/yeyudekuangxiang/imagedesign/internal/errno"
	"github.com/yeyudekuangxiang/imagedesign/internal/validator"
)

func BindForm(c *gin.Context, data interface{}) error {
	if err := c.ShouldBind(data); err != nil {
		err = validator.TranslateError(err)
		return errno.NewBindErr(err)
	}
	return nil
}
