// 非200响应码会正常返回响应，不会返回错误
package resty

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// Execute 使用指定方法请求 http.MethodGet
func Execute(method, url string, hook ReqHook) (*resty.Response, error) {
	if Client == nil {
		Init()
	}
	req := Client.R()
	if hook != nil {
		hook(req)
	}

	res, err := req.Execute(method, url)
	return res, err
}

func Get(url string, hook ReqHook) (*resty.Response, error) {
	if Client == nil {
		Init()
	}
	req := Client.R()
	if hook != nil {
		hook(req)
	}

	res, err := req.Get(url)
	return res, err
}

func Post(url string, hook ReqHook) (*resty.Response, error) {
	if Client == nil {
		Init()
	}
	req := Client.R()
	if hook != nil {
		hook(req)
	}

	res, err := req.Post(url)
	return res, err
}

// PostNew post请求并记录日志 新创建客户端
func PostNew(
	headers map[string]string,
	body, result any,
	url string,
	info map[string]any,
) (*resty.Response, error) {
	req := resty.New().R().SetHeaders(headers).SetBody(body).SetResult(result)
	return PostLog(req, url, info)
}

// PostLog post请求并记录日志 传入请求对象
// 请求前 路径 请求体 自定义日志
// 请求后 路径 请求体 自定义日志 响应码+描述 响应体 耗时
func PostLog(req *resty.Request, url string, info map[string]any) (*resty.Response, error) {
	logger := logrus.WithFields(logrus.Fields{
		"url":     url,
		"request": fmt.Sprintf("%+v", req.Body),
		"info":    info,
	})

	logger.Infoln("post request")
	res, err := req.Post(url)
	if err != nil {
		logger.Errorln("post error", err)
		return nil, err
	}
	logger.WithFields(logrus.Fields{
		"status":   res.Status(),
		"response": res.String(),
		"elapse":   res.Time(),
	}).Infoln("post response")
	return res, nil
}

// PostFile 上传文件，直接读取文件字节
// Content-Type 根据内容检测
func PostFile(url string, hook ReqHook, name string) (*resty.Response, error) {
	bytes, err := os.ReadFile(name)
	if err != nil {
		logrus.Error("read file error ", err)
		return nil, err
	}
	if Client == nil {
		Init()
	}
	req := Client.R().SetBody(bytes)
	if hook != nil {
		hook(req)
	}

	res, err := req.Post(url)
	return res, err
}
