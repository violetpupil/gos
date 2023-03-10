// 语音转写
// 必须先 SetLfAsrSecret 设置语音转写密钥
// https://www.xfyun.cn/doc/asr/ifasr_new/API.html
package xfyun

import (
	"fmt"

	"github.com/sirupsen/logrus"
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
		OrderId          string `json:"orderId"` // 订单id
		TaskEstimateTime int    `json:"taskEstimateTime"`
	} `json:"content"`
}

// Valid 验证语音转写设置是否有效
func (a *xfyun) Valid() error {
	if a == nil || a.appid == "" || a.lfAsrSecret == "" {
		return fmt.Errorf("argument insufficient %+v", a)
	}
	return nil
}

// Upload 上传音频文件
func (a *xfyun) Upload(name string) error {
	err := a.Valid()
	if err != nil {
		return err
	}

	// 查询字符串参数
	info, err := os.Stat(name)
	if err != nil {
		return err
	}
	qs := a.SignA(a.lfAsrSecret)
	qs["fileName"] = info.Name
	qs["fileSize"] = strconv.ToString(info.Size)
	// TODO 音频时长
	qs["duration"] = ""

	res, err := resty.Post(UrlUpload, func(r *resty.Request) *resty.Request {
		r.SetHeader("Content-Type", "application/json")
		r.SetQueryParams(qs)
		return r
	})
	if err != nil {
		return err
	}

	var body UploadRes
	err = Unmarshal(res, &body)
	if err != nil {
		return err
	}
	logrus.Infof("upload success %+v", body)
	return nil
}
