package service

import (
	"bottest/common"
	"bottest/pkg/logger"

	"github.com/tidwall/gjson"
)

//* 此处参考了go-cqhttp内部的部分实现

type ResultGetter interface {
	Get(string) gjson.Result
}

//* ---end---//
// BotReceivePost 消息上报
func BotReceivePost(getter ResultGetter) common.WebError {
	//post_type 上报类型
	switch getter.Get("post_type").Str {
	default:
		//其他
		logger.Warnf("unknow post_type\n")
		return common.ErrInvalidParam().AddMsg(":unknow post_type")
	case "message":
		return BotReceiveMessage(getter)
	case "notice":
		return BotReceiveNotice(getter)
	case "request":
		return BotReceiveRequest(getter)
	}
}

// BotReceiveMessage 消息
func BotReceiveMessage(getter ResultGetter) common.WebError {
	//消息
	//message_type 消息类型
	logger.Infof("message_type: %v\n", getter.Get("message_type").Str)
	switch getter.Get("message_type").Str {
	default:
		//其他
		logger.Warnf("unknow message_type\n")
		return common.ErrInvalidParam().AddMsg(":unknow message_type")
	case "private":
		//私聊
		userID := getter.Get("user_id").Int()
		message := getter.Get("message").Str
		rawMessage := getter.Get("raw_message").Str
		_ = rawMessage
		logger.Infof("Got a private message:\n")
		logger.Infof("from:%v", userID)
		logger.Infof("message:%v", message)
		if userID == 535310511 {
			resMsg := "我已经收到你的消息了，内容是：{ " + message + " }"
			err := BotSendMessage(userID, resMsg, "private")
			if err != nil {
				logger.Warnf("send message err :to:%d ,msg:%s", userID, resMsg)
			}
		}
	case "group":
		//群聊
		logger.Infof("Got a group message:")
		logger.Infof("group:%v", getter.Get("group_id"))
		logger.Infof("from:%v", getter.Get("user_id"))
		logger.Infof("message:%v", getter.Get("message"))

	}

	return nil
}

// BotReceiveNotice 通知
func BotReceiveNotice(getter ResultGetter) common.WebError {
	//通知
	logger.Infof("Got a notice")
	return nil
}

// BotReceiveRequest 请求
func BotReceiveRequest(getter ResultGetter) common.WebError {
	//请求
	logger.Infof("Got a request")
	return nil
}
