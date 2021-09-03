package service

import (
	"bottest/pkg/convert"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var OtherAnswer = []string{
	"什么意思呀？",
	"听不懂听不懂",
	"不听不听王八念经",
	"这就是人类的语言吗，好难懂",
	"今天天气真好~",
	"emmm,要不换个话题吧诶嘿~",
	"不要总说些难以理解的话啦！",
	"这多少有些超出了我的理解范围",
}

var HelloAnswer = []string{
	"好！",
	"你好呀",
	"你好你好",
	"哈喽哈喽",
	"诶嘿~",
}

func TimeCheck() (string, string) {
	now := time.Now().Local()
	h := now.Hour()
	str := convert.ToString(now)
	if h < 8 {
		return "早上", str
	}
	if h < 12 {
		return "上午", str
	}
	if h < 14 {
		return "中午", str
	}
	if h < 19 {
		return "下午", str
	}
	return "晚上", str

}

func GetRandomAnswer(msg string) (bool, string) {
	r := rand.Intn(len(OtherAnswer))
	return true, OtherAnswer[r]
}

func GetHelloAnswer(msg string) (bool, string) {
	h, _ := TimeCheck()
	r := rand.Intn(len(HelloAnswer))

	if r == 0 {
		h = h + HelloAnswer[r]
	} else {
		h = HelloAnswer[r]
	}

	var re = regexp.MustCompile(`(?m)(.安)|(你好)|(泥嚎)|(早上好)|(.午好)|(晚上好)|(晚好)|(嗨)|(哈咯)|(哈喽)|(hi)|(hello)|(.哈哟)`)
	if len(re.FindStringIndex(msg)) > 0 {
		return true, h
	}
	return false, ""
}

// GetAutoAnswer 匹配用/add命令添加的自动回复，支持关键词触发
func GetAutoAnswer(msg string) (bool, string) {
	//todo 2021-09-02 17:00:22 hxx 这里是匹配用/add命令添加的自动回复
	if strings.Contains(msg, "/切割刀防守打法") {
		return true, ""
	}
	return false, ""
}
