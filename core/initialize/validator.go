package initialize

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/yeyudekuangxiang/imagedesign/internal/validator"
)

func InitValidator() {
	binding.Validator = validator.NewValidator()
}
