package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"github.com/yeyudekuangxiang/imagedesign/core/app"
	"github.com/yeyudekuangxiang/imagedesign/internal/util"
	"log"
)

func Middleware(middleware *gin.Engine) {
	middleware.Use(corsM())
	middleware.Use(gin.Recovery())
}

type ThrottleConfig struct {
	Throttle string
}

func throttle() gin.HandlerFunc {
	throttleConfig := &struct {
		Throttle string
	}{}
	_ = app.Ini.Section("http").MapTo(throttleConfig)
	if throttleConfig.Throttle == "" {
		throttleConfig.Throttle = "200-M"
	}
	rate, err := limiter.NewRateFromFormatted(throttleConfig.Throttle)
	if err != nil {
		log.Fatal(err)
	}

	store := memory.NewStoreWithOptions(limiter.StoreOptions{
		Prefix: "throttle",
	})

	middleware := mgin.NewMiddleware(limiter.New(store, rate), mgin.WithKeyGetter(func(c *gin.Context) string {
		return util.Md5(c.ClientIP() + c.Request.Method + c.FullPath())
	}))
	return middleware
}
func corsM() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("x-token", "token", "authorization")
	return cors.New(config)
}
