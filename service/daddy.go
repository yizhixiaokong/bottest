package service

import (
	"bottest/pkg/convert"
	"os"
)

// IsDaddy 判断是否是管理员
func IsDaddy(userID int64) bool {
	daddyUser := os.Getenv("DADDY_USER")
	if daddyUser != "" && userID != convert.ToInt64(daddyUser) {
		return false
	}
	return true
}
