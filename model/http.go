package model

import (
	"boottest/pkg/logger"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

//* 此处参考了go-cqhttp内部的部分实现

type HttpContext struct {
	Ctx *gin.Context
}

func (h HttpContext) Get(k string) gjson.Result {
	c := h.Ctx
	if q := c.Query(k); q != "" {
		return gjson.Result{Type: gjson.String, Str: q}
	}
	if c.Request.Method == "POST" {
		if h := c.Request.Header.Get("Content-Type"); h != "" {

			if strings.Contains(h, "application/x-www-form-urlencoded") {
				if p, ok := c.GetPostForm(k); ok {
					return gjson.Result{Type: gjson.String, Str: p}
				}
			}
			if strings.Contains(h, "application/json") {
				if obj, ok := c.Get("json_body"); ok {
					return obj.(gjson.Result).Get(k)
				} else {
					d, err := c.GetRawData()
					if err != nil {
						logger.Warnf("获取请求 %v 的Body时出现错误: %v", c.Request.RequestURI, err)
						c.Status(400)
						return gjson.Result{Type: gjson.Null, Str: ""}
					}
					c.Set("json_body", gjson.ParseBytes(d))
					return gjson.ParseBytes(d).Get(k)
				}
			}
		}
	}
	return gjson.Result{Type: gjson.Null, Str: ""}
}

//* ---end---//
