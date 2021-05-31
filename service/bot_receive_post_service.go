package service

import (
	"bottest/common"
	"bottest/pkg/logger"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/tidwall/gjson"
)

//* 此处参考了go-cqhttp内部的部分实现

type ResultGetter interface {
	Get(string) gjson.Result
}

//* ---end---//

func BotReceivePost(getter ResultGetter) common.WebError {
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
		logger.Infof("Got a private message:\n")
		logger.Infof("from:%v", userID)
		logger.Infof("message:%v", message)
		if userID == 535310511 {
			resMsg := make(map[string]interface{})
			resMsg["user_id"] = userID
			resMsg["message"] = "我已经收到你的消息了，内容是：{ " + message + " }"
			client := &http.Client{}
			url := "http://127.0.0.1:5700/send_private_msg"
			req, _ := json.Marshal(resMsg)
			// logger.Infof("%v", string(req))
			req_new := bytes.NewBuffer([]byte(req))
			request, _ := http.NewRequest("POST", url, req_new)
			request.Header.Set("Content-type", "application/json")
			response, _ := client.Do(request)
			logger.Infof("%v", response)
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

func BotReceiveNotice(getter ResultGetter) common.WebError {
	//通知
	logger.Infof("Got a notice")
	return nil
}
func BotReceiveRequest(getter ResultGetter) common.WebError {
	//请求
	logger.Infof("Got a request")
	return nil
}
