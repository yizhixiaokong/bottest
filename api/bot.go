package api

import (
	"boottest/common"
	"boottest/model"
	"boottest/pkg/logger"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BotMessage 接收bot消息接口
func BotMessage(c *gin.Context) {

	getter := model.HttpContext{Ctx: c}

	//post_type 上报类型
	logger.Infof("post_type Object: %v\n", getter.Get("post_type"))
	pt := getter.Get("post_type").Str
	logger.Infof("post_type: %v\n", pt)
	switch pt {
	case "message":
		//消息
		//message_type 消息类型
		logger.Infof("message_type: %v\n", getter.Get("message_type").Str)
		switch getter.Get("message_type").String() {
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
		default:
			//其他
			logger.Warnf("unknow message_type\n")
			common.ResJson(c, "unknow message_type", nil)
			return
		}
	case "notice":
		//通知
		logger.Infof("Got a notice")
	case "request":
		//请求
		logger.Infof("Got a request")
	default:
		//其他
		logger.Warnf("unknow post_type\n")
		common.ResJson(c, "unknow post_type", nil)
		return
	}

	// logger.Infof("%+v\n", service)
	common.ResJson(c, "got it", nil)
}

func GetPostType(getter model.ResultGetter) {

}
