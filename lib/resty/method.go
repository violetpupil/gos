// 非200响应码会正常返回响应，不会返回错误
package resty

import (
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
