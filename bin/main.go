package main

import (
	"flag"
	"github.com/dunpju/higo-gin/higo"
	"github.com/dunpju/higo-logger/logger"
	"github.com/dunpju/higo-utils/utils"
	"higo-framework/app/beans"
	"higo-framework/app/middlewares"
	"higo-framework/router"
	"strings"
)

// 定义命令行参数
var Root string

func main() {
	flag.StringVar(&Root, "root", "."+utils.Dir.Separator(), "主目录")

	//解析命令行参数
	//flag.Parse()

	logger.Logrus.IsInit(false).Init()
	logger.Logrus.Info("root:", Root)

	higo.RecoverHandle = middlewares.RecoverHandle
	higo.Init(utils.Slice.New(strings.Split(Root, utils.Dir.Separator())...)).
		CorsHandlerFunc(middlewares.NewCors()).
		Middleware(middlewares.NewRunLog()).
		AddServe(router.NewHttps()).
		Beans(beans.NewBean()).
		Boot()
}
