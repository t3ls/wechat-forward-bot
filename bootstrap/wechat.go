package bootstrap

import (
	"wechat-forward-bot/config"
	"wechat-forward-bot/handler/wechat"

	"github.com/eatmoreapple/openwechat"
	log "github.com/sirupsen/logrus"
)

func StartWebChat() {
	log.Info("Start WebChat Bot")
	bot := openwechat.DefaultBot(openwechat.Desktop)
	bot.MessageHandler = wechat.Handler.AsMessageHandler()
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	err := bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption())
	if err != nil {
		log.Fatal(err)
		return
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		log.Fatal(err)
		return
	}

	targetUser := config.GetForwardTargetUsername()
	if targetUser == "" {
		log.Fatal("No target user specified")
		return
	}

	friends, err := self.Friends()
	for i, friend := range friends {
		log.Println(i, friend)
		if friend.NickName == targetUser {
			log.Printf("Found target user: %v", friend)
			wechat.TargetUser = friend
			_, err = friend.SendText("你好，我是机器人，我已经准备好转发所有收到的消息给您了。")
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}

	groups, err := self.Groups()
	for i, group := range groups {
		log.Println(i, group)
	}

	err = bot.Block()
	if err != nil {
		log.Fatal(err)
		return
	}
}
