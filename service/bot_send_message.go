package service

import (
	"bottest/common"
	"bottest/pkg/logger"
	"bottest/util"
)

const (
	sendIP string = "127.0.0.1:5700"

	SendPrivateMsg string = "/send_private_msg"
	SendGroupMsg   string = "/send_group_msg"
	SendMsg        string = "/send_msg"
)

// BotSendMessage 发消息
func BotSendMessage(userID int64, msg string, msgType string) common.WebError {

	switch msgType {
	default:
		//其他
		logger.Warnf("unknow message_type\n")
		return common.ErrInvalidParam().AddMsg(":unknow message_type")
	case "private":
		// 发私聊
		return BotSendPrivateMessage(userID, msg)
	case "group":
		// 发群聊
		return BotSendGroupMessage(userID, msg)
	}

}

// BotSendPrivateMessage 私聊
func BotSendPrivateMessage(userID int64, msg string) common.WebError {
	reqMsg := make(map[string]interface{})
	reqMsg["user_id"] = userID
	reqMsg["message"] = msg
	reqMsg["auto_escape"] = false

	res, err := util.SendHttpPost(sendIP, SendPrivateMsg, reqMsg)
	if err != nil {
		logger.Errorf("BotSendPrivateMessage :SendHttpPost err:%v", err)
		return common.ErrServer()
	}
	if res.Get("status").Str != "ok" {
		logger.Errorf("BotSendPrivateMessage err: res msg:%s", res.Get("msg").Str)
		return common.ErrExecFaild()
	}
	return nil
}

// BotSendGroupMessage 群聊
func BotSendGroupMessage(groupID int64, msg string) common.WebError {
	reqMsg := make(map[string]interface{})
	reqMsg["group_id"] = groupID
	reqMsg["message"] = msg

	res, err := util.SendHttpPost(sendIP, SendGroupMsg, reqMsg)
	if err != nil {
		logger.Errorf("BotSendPrivateMessage :SendHttpPost err:%v", err)
		return common.ErrServer()
	}
	if res.Get("status").Str != "ok" {
		logger.Errorf("BotSendPrivateMessage err: res msg:%s", res.Get("msg").Str)
		return common.ErrExecFaild()
	}
	return nil
}
