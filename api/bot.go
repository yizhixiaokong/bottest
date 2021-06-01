package api

import (
	"bottest/common"
	"bottest/util"

	"bottest/service"

	"github.com/gin-gonic/gin"
)

// BotMessage 接收bot消息接口
func BotMessage(c *gin.Context) {
	err := service.BotReceivePost(util.HttpContext{Ctx: c})
	common.ResJson(c, nil, err)
	return
}
