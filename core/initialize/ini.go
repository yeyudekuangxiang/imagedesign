package initialize

import (
	"github.com/yeyudekuangxiang/imagedesign/core/app"
	"gopkg.in/ini.v1"
	"log"
)

func InitIni(source interface{}) {
	f, err := ini.Load(source)
	if err != nil {
		log.Fatal(err)
	}
	app.Ini = f
}
