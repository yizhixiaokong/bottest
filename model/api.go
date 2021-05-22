package model

import "github.com/tidwall/gjson"

//* 此处参考了go-cqhttp内部的部分实现

type ResultGetter interface {
	Get(string) gjson.Result
}

//* ---end---//
