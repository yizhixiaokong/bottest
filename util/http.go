package util

import (
	"bottest/common"
	"bottest/pkg/convert"
	"bottest/pkg/logger"

	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

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

var httpPrefix = "http://"

func SendHttpPost(serverIP string, url string, reqData interface{}) (gjson.Result, error) {

	if !strings.HasPrefix(url, httpPrefix) {
		url = httpPrefix + serverIP + url
	}
	contentType := "application/json; charset=utf-8"

	client := &http.Client{
		Timeout: 8 * time.Second,
	}

	requestJson, err := json.Marshal(reqData)
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}

	logger.Infof("method: POST url: %v", url)
	logger.Infof("request body: %v", string(requestJson))
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestJson))
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}
	request.Header.Set("Content-Type", contentType)

	var resp *http.Response
	resp, err = client.Do(request)
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}

	// 反序列化
	var respJson []byte
	respJson, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}

	logger.Debugf("response body: %v", string(respJson))
	res := gjson.ParseBytes(respJson)

	return res, nil
}

func SendHttpPut(serverIP string, url string, reqData interface{}, respData interface{}) (gjson.Result, error) {

	var respModel common.ResponseModel
	respModel.ErrorCode = -1

	if !strings.HasPrefix(url, httpPrefix) {
		url = httpPrefix + serverIP + url
	}
	contentType := "application/json; charset=utf-8"

	client := &http.Client{
		Timeout: 8 * time.Second,
	}

	requestJson, err := json.Marshal(reqData)
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}

	logger.Infof("method: PUT url: %v", url)
	logger.Infof("request body: %v", string(requestJson))
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(requestJson))
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}
	request.Header.Set("Content-Type", contentType)

	var resp *http.Response
	resp, err = client.Do(request)
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}

	// 反序列化
	var respJson []byte
	respJson, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}

	logger.Debugf("response body: %v", string(respJson))
	res := gjson.ParseBytes(respJson)

	return res, nil
}

func SendHttpGet(serverIP string, url string, reqData map[string]string, respData interface{}) (gjson.Result, error) {

	var respModel common.ResponseModel
	respModel.ErrorCode = -1

	if !strings.HasPrefix(url, httpPrefix) {
		url = httpPrefix + serverIP + url
	}

	client := &http.Client{
		Timeout: 8 * time.Second,
	}

	if len(reqData) > 0 {
		url += "?"
		for k, v := range reqData {
			url += fmt.Sprintf("%s=%s&", k, v)
		}
		url = url[:len(url)-1]
	}

	// 发起请求
	logger.Debugf("method: GET url: %v", url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}

	var resp *http.Response
	resp, err = client.Do(request)
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}

	// 反序列化
	var respJson []byte
	respJson, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}

	logger.Debugf("response body: %v", string(respJson))
	res := gjson.ParseBytes(respJson)

	return res, nil
}

func SendHttpDelete(serverIP string, url string, reqData map[string]string, respData interface{}) (gjson.Result, error) {

	var respModel common.ResponseModel
	respModel.ErrorCode = -1

	if !strings.HasPrefix(url, httpPrefix) {
		url = httpPrefix + serverIP + url
	}

	client := &http.Client{
		Timeout: 8 * time.Second,
	}

	if len(reqData) > 0 {
		url += "?"
		for k, v := range reqData {
			url += fmt.Sprintf("%s=%s&", k, v)
		}
		url = url[:len(url)-1]
	}

	// 发起请求
	logger.Debugf("method: Delete url: %v", url)

	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}

	var resp *http.Response
	resp, err = client.Do(request)
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}

	// 反序列化
	var respJson []byte
	respJson, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}

	logger.Debugf("response body: %v", string(respJson))
	res := gjson.ParseBytes(respJson)

	return res, nil
}

func SendHttpDeleteJson(serverIP string, url string, reqData interface{}, respData interface{}) (gjson.Result, error) {
	var respModel common.ResponseModel
	respModel.ErrorCode = -1

	if !strings.HasPrefix(url, httpPrefix) {
		url = httpPrefix + serverIP + url
	}
	contentType := "application/json; charset=utf-8"

	client := &http.Client{
		Timeout: 8 * time.Second,
	}

	requestJson, err := json.Marshal(reqData)
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}

	logger.Infof("method: DELETE url: %v", url)
	logger.Infof("request body: %v", string(requestJson))
	request, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(requestJson))
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}
	request.Header.Set("Content-Type", contentType)

	var resp *http.Response
	resp, err = client.Do(request)
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}

	// 反序列化
	var respJson []byte
	respJson, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		return gjson.Result{Type: gjson.Null, Str: ""}, err
	}

	logger.Debugf("response body: %v", string(respJson))
	res := gjson.ParseBytes(respJson)

	return res, nil
}

func Resp2struct(respBody io.Reader, respStruct interface{}) error {
	var respJson []byte
	var err error

	respJson, err = ioutil.ReadAll(respBody)
	if err != nil {
		return err
	}

	logger.Debugf("response body: %v", string(respJson))

	var respModel common.ResponseModel
	if err = json.Unmarshal(respJson, &respModel); err != nil {
		logger.Error(err.Error())
		return err
	}

	if respModel.ErrorCode != 0 {
		return fmt.Errorf("err code: %v, err msg: %v", respModel.ErrorCode, respModel.Message)
	}

	var resDataJson []byte
	resDataJson, err = json.Marshal(respModel.Data)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = json.Unmarshal(resDataJson, respStruct)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func Interface2struct(data interface{}, result interface{}) error {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = json.Unmarshal(jsonByte, result)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return err
}

func Struct2Map(data interface{}) map[string]string {
	m := make(map[string]interface{})
	jsonByte, _ := json.Marshal(data)

	_ = json.Unmarshal(jsonByte, &m)

	res := make(map[string]string)
	for k, v := range m {
		res[k] = convert.ToString(v)
	}

	return res
}
