package service

import "math/rand"

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

func RandomAnswer(s string) (bool, string) {
	r := rand.Intn(len(OtherAnswer))
	return true, OtherAnswer[r]
}
