// 科大讯飞api调用
// 必须先 Init 初始化
// 不同服务的密钥不同
// https://www.xfyun.cn/doc/
package xfyun

import (
	"fmt"

	"github.com/violetpupil/components/crypto/sign"
	"github.com/violetpupil/components/std/time"
)

type xfyun struct {
	appid string
}

var Xfyun *xfyun

// Init 初始化api调用
func Init(appid string) {
	Xfyun = &xfyun{appid}
}

// SignA 生成签名并构造请求参数，secret为服务密钥
func (a *xfyun) SignA(secret string) map[string]string {
	ts := time.Ts()
	signA := sign.Sign(a.appid, ts, secret)
	m := map[string]string{
		"signa": signA,
		"appId": a.appid,
		"ts":    ts,
	}
	return m
}

// 处理成功
const ResCodeSuss = "000000"

// ResBody 通用响应体
type ResBody struct {
	Code     string `json:"code"`     // 处理码
	DescInfo string `json:"descInfo"` // 处理结果
}

// Error 处理失败时，获取错误
func (r ResBody) Error() error {
	return fmt.Errorf("xfyun api: %s %s", r.Code, r.DescInfo)
}
