package main

import (
	"wechat-forward-bot/bootstrap"
	"wechat-forward-bot/config"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)

	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	log.Info("程序启动")
	err := config.LoadConfig()
	if err != nil {
		log.Warn("没有找到配置文件，尝试读取环境变量")
	}

	bootstrap.StartWebChat()

	log.Info("程序退出")
}
