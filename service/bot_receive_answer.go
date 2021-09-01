package service

import "math/rand"

var OtherAnswer = []string{
	"什么意思呀？",
	"听不懂听不懂",
	"这就是人类的语言吗，好难懂",
	"你要不说点别的？",
	"不要总说些难以理解的话啦！",
}

func RandomAnswer() string {
	r := rand.Intn(len(OtherAnswer))
	return OtherAnswer[r]
}
