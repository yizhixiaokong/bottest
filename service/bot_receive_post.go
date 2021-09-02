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
		// logger.Infof("Got a private message:")
		// logger.Infof("from:%v", userID)
		// logger.Infof("message:%v", message)
		logger.Infof("Got a private message from user:%v message:%v", userID, message)

		var resMsg string

		//todo 2021-09-01 14:29:10 hxx 目前只有管理员用户开放私聊功能，之后可改成添加好友后开放
		if !IsDaddy(userID) {
			resMsg = "妈妈说不可以和陌生人说话的"
		} else {
			resMsg = "我已经收到你的消息了，内容是：{ " + message + " }\n"
			resMsg += BotReceiveMessageInfo(message, false)
		}
		err := BotSendMessage(userID, resMsg, "private")
		if err != nil {
			logger.Warnf("send message err :to:%d ,msg:%s", userID, resMsg)
		}

	case "group":
		//群聊
		// logger.Infof("Got a group message:")

		groupID := getter.Get("group_id").Int()
		userID := getter.Get("user_id").Int()
		message := getter.Get("message").Str
		rawMessage := getter.Get("raw_message").Str
		_ = rawMessage
		logger.Infof("Got a group message from group:%v user:%v message:%v", groupID, userID, message)
		// logger.Infof("group:%v", groupID)
		// logger.Infof("from:%v", userID)
		// logger.Infof("message:%v", message)
		// logger.Infof("raw_message:%v", rawMessage)

		// logger.Infof("msg:%v", msg)
		resMsg := BotReceiveMessageInfo(message, true)
		if resMsg != "" {
			// logger.Infof("resMsg:%v", resMsg)
			err := BotSendMessage(groupID, resMsg, "group")
			if err != nil {
				logger.Warnf("send message err :to:%d ,msg:%s", groupID, resMsg)
			}

		}
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

func BotReceiveMessageInfo(msg string, group bool) string {

	//检测AT是否符合要求
	checkAt := func(f func(str string) (bool, string), needAt bool) (bool, string) {
		if !needAt {
			return f(msg)
		}
		isAt, m := IsAt(ReturnSelf(), msg)
		if isAt {
			return f(m)
		}
		return false, ""
	}

	//todo 暂定回复消息分三个模块：命令、功能、自动回复。

	//todo暂定命令只能管理员用户使用
	// 以下按排列顺序判定触发优先级
	// ——————————————————命令——————————————————
	// todo 2021-09-01 14:33:35 hxx 增加对话注册命令，注册的对话保存数据库，暂定所有群通用
	// #add 添加自动回复
	// #闭嘴 不再发送消息
	// #说话 可以继续发送消息

	// ——————————————————功能——————————————————
	// /help 返回功能菜单
	if if_help, res := checkAt(HelpMenu, false); if_help {
		return res
	}
	// /random n 返回0~n的一个数
	if if_random, res := checkAt(Random, false); if_random {
		return res
	}

	// ——————————————————自动回复——————————————————
	// 用/add命令添加的自动回复，支持部分匹配触发

	// ——————————————————其他回复——————————————————
	_, str := checkAt(RandomAnswer, true)
	return str
}
