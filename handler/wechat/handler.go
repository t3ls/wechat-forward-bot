package wechat

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"wechat-forward-bot/config"
)

var Handler = openwechat.NewMessageMatchDispatcher()
var TargetUser *openwechat.Friend

func init() {
	Handler.SetAsync(true)
	Handler.OnFriend(RawFriendMessageHandler)
	Handler.OnGroup(RawGroupMessageHandler)
}

func RawFriendMessageHandler(ctx *openwechat.MessageContext) {
	msg := ctx.Message

	sender, err := msg.Sender()
	if err != nil {
		log.Errorf("Failed to get sender: %v", err)
		return
	}
	log.Debugf("Received Text Msg From %v: %v", sender.NickName, msg.Content)

	// Forward the message to the target user
	keyword := config.GetWechatKeyword()
	if keyword != "" {
		if msg.IsText() && !strings.Contains(msg.Content, keyword) {
			log.Debugf("Ignore message from %v: %v", sender.NickName, msg.Content)
			return
		}
	}

	if TargetUser == nil {
		log.Fatal("No target user specified")
		return
	}

	if msg.IsPicture() {
		log.Info("Received Image Msg, saving to cache")
		resp, err := msg.GetFile()
		if err != nil {
			log.Errorf("Failed to get file: %v", err)
			return
		}
		defer resp.Body.Close()
		_, err = TargetUser.SendText(fmt.Sprintf("收到【图片】消息: 【发送者】昵称：%s, 备注：%s", sender.NickName, sender.RemarkName))
		if err != nil {
			log.Errorf("Failed to send text: %v", err)
			return
		}
		_, err = TargetUser.SendImage(resp.Body)
		if err != nil {
			log.Errorf("Failed to send image: %v", err)
			return
		}
	} else if msg.IsVoice() {
		log.Info("Received Voice Msg, saving to cache")
		resp, err := msg.GetVoice()
		if err != nil {
			log.Errorf("Failed to get voice: %v", err)
			return
		}
		defer resp.Body.Close()
		file, err := os.CreateTemp("", "wechat_handle.voice.*.mp3")
		if err != nil {
			log.Errorf("Failed to create temp file: %v", err)
			return
		}
		defer file.Close()
		defer os.Remove(file.Name())
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Errorf("Failed to copy file: %v", err)
			return
		}
		file.Seek(0, 0)
		_, err = TargetUser.SendText(fmt.Sprintf("收到【语音】消息: 【发送者】昵称：%s, 备注：%s 【语音长度】%f秒", sender.NickName, sender.RemarkName, float32(msg.VoiceLength)/1000))
		if err != nil {
			log.Errorf("Failed to send text: %v", err)
			return
		}
		_, err = TargetUser.SendFile(file)
		if err != nil {
			log.Errorf("Failed to send file: %v", err)
			return
		}
	} else if msg.IsFriendAdd() {
		_, err = TargetUser.SendText(fmt.Sprintf("收到【好友添加】消息: 【发送者】昵称：%s, 备注：%s", sender.NickName, sender.RemarkName))
		if err != nil {
			log.Errorf("Failed to send text: %v", err)
			return
		}
	} else if msg.IsVideo() {
		log.Info("Received Video Msg, saving to cache")
		resp, err := msg.GetFile()
		if err != nil {
			log.Errorf("Failed to get file: %v", err)
			return
		}
		defer resp.Body.Close()
		file, err := os.CreateTemp("", "wechat_handle.video.*.mp4")
		if err != nil {
			log.Errorf("Failed to create temp file: %v", err)
			return
		}
		defer file.Close()
		defer os.Remove(file.Name())
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Errorf("Failed to copy file: %v", err)
			return
		}
		file.Seek(0, 0)
		_, err = TargetUser.SendText(fmt.Sprintf("收到【视频】消息: 【发送者】昵称：%s, 备注：%s", sender.NickName, sender.RemarkName))
		if err != nil {
			log.Errorf("Failed to send text: %v", err)
			return
		}
		_, err = TargetUser.SendVideo(file)
		if err != nil {
			log.Errorf("Failed to send video: %v", err)
			return
		}
	} else if msg.IsVideo() {
		log.Info("Received Video Msg, saving to cache")
		resp, err := msg.GetFile()
		if err != nil {
			log.Errorf("Failed to get file: %v", err)
			return
		}
		defer resp.Body.Close()
		file, err := os.CreateTemp("", "wechat_handle.video.*.mp4")
		if err != nil {
			log.Errorf("Failed to create temp file: %v", err)
			return
		}
		defer file.Close()
		defer os.Remove(file.Name())
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Errorf("Failed to copy file: %v", err)
			return
		}
		file.Seek(0, 0)
		_, err = TargetUser.SendText(fmt.Sprintf("收到【视频】消息: \n【发送者】昵称：%s, 备注：%s", sender.NickName, sender.RemarkName))
		if err != nil {
			log.Errorf("Failed to send text: %v", err)
			return
		}
		_, err = TargetUser.SendVideo(file)
		if err != nil {
			log.Errorf("Failed to send video: %v", err)
			return
		}
	} else if msg.IsEmoticon() {
		_, err = TargetUser.SendText(fmt.Sprintf("收到【表情】消息: \n【发送者】昵称：%s, 备注：%s", sender.NickName, sender.RemarkName))
		if err != nil {
			log.Errorf("Failed to send text: %v", err)
			return
		}
	} else if msg.IsFriendAdd() {
		_, err = TargetUser.SendText(fmt.Sprintf("收到【好友添加】消息: \n【发送者】昵称：%s, 备注：%s", sender.NickName, sender.RemarkName))
		if err != nil {
			log.Errorf("Failed to send text: %v", err)
			return
		}
	} else if msg.IsLocation() {
		_, err = TargetUser.SendText(fmt.Sprintf("收到【位置】消息: \n【发送者】昵称：%s, 备注：%s", sender.NickName, sender.RemarkName))
		if err != nil {
			log.Errorf("Failed to send text: %v", err)
			return
		}
	} else if msg.IsCard() {
		_, err = TargetUser.SendText(fmt.Sprintf("收到【名片】消息: \n【发送者】昵称：%s, 备注：%s", sender.NickName, sender.RemarkName))
		if err != nil {
			log.Errorf("Failed to send text: %v", err)
			return
		}
	} else if msg.IsPaiYiPai() {
		_, err = TargetUser.SendText(fmt.Sprintf("收到【拍一拍】: \n【发送者】昵称：%s, 备注：%s", sender.NickName, sender.RemarkName))
		if err != nil {
			log.Errorf("Failed to send text: %v", err)
			return
		}
	} else if msg.IsText() {
		_, err = TargetUser.SendText(fmt.Sprintf("收到【文本】消息: \n【发送者】昵称：%s, 备注：%s \n【消息内容】%s", sender.NickName, sender.RemarkName, msg.Content))
		if err != nil {
			log.Errorf("Failed to send text: %v", err)
			return
		}
	}
}

func RawGroupMessageHandler(ctx *openwechat.MessageContext) {
	msg := ctx.Message

	sender, err := msg.Sender()
	if err != nil {
		log.Errorf("Failed to get sender: %v", err)
		return
	}
	log.Debugf("Received Text Msg From %v: %v", sender.NickName, msg.Content)
	if msg.IsText() {
		_, err = TargetUser.SendText(fmt.Sprintf("收到【群聊文本】消息: \n【发送者】昵称：%s, 备注：%s \n【消息内容】%s", sender.NickName, sender.RemarkName, msg.Content))
		if err != nil {
			log.Errorf("Failed to send text: %v", err)
			return
		}
	}
	// TODO
}
