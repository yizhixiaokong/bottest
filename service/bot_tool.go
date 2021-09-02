package service

import (
	"bottest/pkg/convert"
	"os"
	"strings"
)

func ReturnDaddy() int64 {
	daddyUser := os.Getenv("DADDY_USER")
	return convert.ToInt64(daddyUser)
}

func ReturnSelf() int64 {
	selfUser := os.Getenv("SELF_USER")
	return convert.ToInt64(selfUser)
}

// IsDaddy 判断是否是管理员
func IsDaddy(userID int64) bool {
	daddyUser := os.Getenv("DADDY_USER")
	if daddyUser != "" && userID != convert.ToInt64(daddyUser) {
		return false
	}
	return true
}

// IsAt 判断是否被艾特
func IsAt(userID int64, msg string) (bool, string) {
	// logger.Infof("userID :%v", userID)
	// logger.Infof("msg :%v", msg)
	at := "[CQ:at,qq=" + convert.ToString(userID) + "]"
	atLen := len(at)
	index := strings.Index(msg, at)
	// logger.Infof("index:%v", index)
	if index >= 0 {
		return true, string([]byte(msg)[:index]) + string([]byte(msg)[index+atLen:])
	}
	return false, msg
}
