package main

import (
	"flag"
	"github.com/dengpju/higo-gin/higo"
	"github.com/dengpju/higo-logger/logger"
	"github.com/dengpju/higo-utils/utils"
	"higo-framework/app/beans"
	"higo-framework/app/middlewares"
	"higo-framework/router"
	"strings"
)

//定义命令行参数
var Root string
var AlhttpsLogPath string

func main()  {
	flag.StringVar(&Root, "root", "."+utils.PathSeparator(), "主目录")
	flag.StringVar(&AlhttpsLogPath, "log", "."+utils.PathSeparator()+"log", "alhttps日志路径")

	//解析命令行参数
	//flag.Parse()

	logger.Logrus.IsInit(false).Init()
	logger.Logrus.Info("root:", Root)
	logger.Logrus.Info("alhttps log:", AlhttpsLogPath)

	beanConfig := beans.NewMyBean()
	higo.Init(utils.NewSliceString(strings.Split(Root, utils.PathSeparator())...)).
		AuthHandlerFunc(middlewares.NewAuth()).
		Middleware(middlewares.NewRunLog()).
		AddServe(router.NewHttps(), beanConfig).
		IsAutoTLS(true).
		IsRedisPool().
		InitGroupIsAuth(true).
		Cron("* * * * * ?", func() {
			utils.SecretExpiredClear()
			logger.Logrus.Infoln("执行密钥清除,还有", utils.SecretContainer.Len(), "个秘钥")
		}).
		SetBits(1024).
		Boot()
}
