// 语音转写
// https://www.xfyun.cn/doc/asr/ifasr_new/API.html
package xfyun

import (
	"github.com/violetpupil/components/lib/resty"
	"github.com/violetpupil/components/std/os"
	"github.com/violetpupil/components/std/strconv"
)

const (
	// 文件上传地址
	UrlUpload = "https://raasr.xfyun.cn/v2/api/upload"
	// 获取结果地址
	UrlGetResult = "https://raasr.xfyun.cn/v2/api/getResult"
)

// UploadRes 上传音频文件响应体
type UploadRes struct {
	ResBody
	Content struct {
		OrderId          string `json:"orderId"`
		TaskEstimateTime int    `json:"taskEstimateTime"`
	} `json:"content"`
}

// Upload 上传音频文件
func (a *xfyun) Upload(secret, name string) error {
	// 查询字符串参数
	info, err := os.Stat(name)
	if err != nil {
		return err
	}
	qs := a.SignA(secret)
	qs["fileName"] = info.Name
	qs["fileSize"] = strconv.ToString(info.Size)
	// TODO 音频时长
	qs["duration"] = ""

	_, err = resty.Post(UrlUpload, func(r *resty.Request) *resty.Request {
		r.SetHeader("Content-Type", "application/json")
		r.SetQueryParams(qs)
		return r
	})
	if err != nil {
		return err
	}

	// TODO
	return nil
}
