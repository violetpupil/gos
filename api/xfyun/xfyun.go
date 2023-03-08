// https://www.xfyun.cn/doc/
package xfyun

import "github.com/violetpupil/components/lib/resty"

// 语音转写
// https://www.xfyun.cn/doc/asr/ifasr_new/API.html
const (
	// 文件上传
	UrlUpload = "https://raasr.xfyun.cn/v2/api/upload"
	// 获取结果
	UrlGetResult = "https://raasr.xfyun.cn/v2/api/getResult"
)

// Upload 上传音频文件
func Upload() {
	resty.Post(UrlUpload, func(r *resty.Request) *resty.Request {
		r.SetHeader("Content-Type", "application/json; charset=UTF-8,Chunked: false")
		return r
	})
}
