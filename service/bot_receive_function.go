package service

import "strings"

var helpStr string = `这里是Shiro酱的帮助菜单，目前的功能有：
1. 发送 /help 显示功能菜单
2. 发送 /random n 返回一个0~n-1的随机数（未实现）`

// HelpMenu /help 显示功能菜单
func HelpMenu(msg string) (bool, string) {
	if strings.Contains(msg, "/help") {
		return true, helpStr
	}
	return false, ""
}
