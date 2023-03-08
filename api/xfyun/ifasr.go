// 语音转写
// https://www.xfyun.cn/doc/asr/ifasr_new/API.html
package xfyun

import (
	"github.com/violetpupil/components/crypto/sign"
	"github.com/violetpupil/components/lib/resty"
	"github.com/violetpupil/components/std/time"
)

const (
	// 文件上传地址
	UrlUpload = "https://raasr.xfyun.cn/v2/api/upload"
	// 获取结果地址
	UrlGetResult = "https://raasr.xfyun.cn/v2/api/getResult"
)

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

// Upload 上传音频文件
func (a *xfyun) Upload(secret string) {
	qs := a.SignA(secret)
	qs["fileName"] = ""
	qs["fileSize"] = ""
	qs["duration"] = ""
	resty.Post(UrlUpload, func(r *resty.Request) *resty.Request {
		r.SetHeader("Content-Type", "application/json; charset=UTF-8,Chunked: false")
		r.SetQueryParams(qs)
		return r
	})
}
