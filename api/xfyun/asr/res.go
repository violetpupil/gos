package asr

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/lib/resty"
)

// 响应成功时，才能获取到响应内容
const ResCodeSuss = "000000"

// ResBody 通用响应体
type ResBody struct {
	Code     string `json:"code"`     // 响应代码
	DescInfo string `json:"descInfo"` // 响应信息
}

// Error 处理失败时，获取错误
func (r ResBody) Error() error {
	return fmt.Errorf("xfyun api: %s %s", r.Code, r.DescInfo)
}

// CodeI 获取响应代码，实现ResBodyI
func (r ResBody) CodeI() string {
	return r.Code
}

// ResBodyI 通用响应体接口
// 嵌入了 ResBody 的结构体可以统一使用该接口
type ResBodyI interface {
	Error() error
	CodeI() string
}

// Unmarshal 响应体json解码，响应代码不是成功时返回错误
// body 必须传不为nil的指针
func Unmarshal(res *resty.Res, body ResBodyI) error {
	err := json.Unmarshal(res.Body, body)
	if err != nil {
		logrus.Error("json unmarshal error ", err)
		return res.ToError()
	}
	// 解码成功，但没有获取到响应代码
	if body.CodeI() == "" {
		return res.ToError()
	} else if body.CodeI() != ResCodeSuss {
		return body.Error()
	}
	return nil
}
