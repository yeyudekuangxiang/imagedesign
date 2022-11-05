package main

import (
	"flag"
	"github.com/yeyudekuangxiang/imagedesign/core/initialize"
	"os"
	"os/signal"
)

var (
	flagConf = flag.String("conf", "./config.ini", "-c")
)

func init() {
	flag.Parse()
	initialize.InitIni(*flagConf)
	initialize.InitValidator()
	initialize.InitServer()
}
func main() {
	initialize.RunServer()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	initialize.CloseServer()
}
