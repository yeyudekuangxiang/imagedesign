package initialize

import (
	"github.com/yeyudekuangxiang/imagedesign/core/app"
	"github.com/yeyudekuangxiang/imagedesign/internal/db"
	"log"
)

func InitDB() {
	var conf db.Config
	err := app.Ini.Section("database").MapTo(&conf)
	if err != nil {
		log.Fatal(err)
	}
	//创建晓筑规范数据库连接
	gormDb, err := db.NewDB(conf)
	if err != nil {
		log.Fatal(err)
	}
	*app.DB = *gormDb
}
