package service

import (
	"bottest/pkg/convert"
	"math/rand"
	"regexp"
	"strings"
)

var helpStr string = `这里是Shiro酱的帮助菜单，目前的功能有：
1. 发送 /help 显示功能菜单
2. 发送 /random n 返回一个0~n的随机数`

// GetHelpMenu /help 显示功能菜单
func GetHelpMenu(msg string) (bool, string) {
	if strings.Contains(msg, "/help") {
		return true, helpStr
	}
	return false, ""
}

// GetRandom /random n 返回0~n的一个数
func GetRandom(msg string) (bool, string) {

	var re = regexp.MustCompile(`(?m)/random (\d+)`)

	// logger.Infof("msg:%v", msg)
	if len(re.FindStringIndex(msg)) > 0 {
		match := re.FindStringSubmatch(msg)
		if len(match) > 1 {
			num := match[1]
			//0~n的随机数
			reStr := "(随机0~" + num + ")：" + convert.ToString(rand.Intn(convert.ToInt(num)+1))
			return true, reStr
		}
	}
	return false, msg
}
