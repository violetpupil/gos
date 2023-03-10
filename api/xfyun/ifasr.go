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
		OrderId          string `json:"orderId"`          // 订单id
		TaskEstimateTime int    `json:"taskEstimateTime"` // 订单预估耗时，单位毫秒
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
func (a *xfyun) Upload(name string) (*UploadRes, error) {
	// 检查密钥设置
	err := a.Valid()
	if err != nil {
		return nil, err
	}

	// 查询字符串参数
	info, err := os.Stat(name)
	if err != nil {
		return nil, err
	}
	qs := a.SignA(a.lfAsrSecret)
	qs["fileName"] = info.Name
	qs["fileSize"] = strconv.ToString(info.Size)
	// TODO 音频时长
	qs["duration"] = ""

	// 请求体
	bytes, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	res, err := resty.Post(UrlUpload, func(r *resty.Request) *resty.Request {
		r.SetHeader("Content-Type", "application/json")
		r.SetQueryParams(qs)
		r.SetBody(bytes)
		return r
	})
	if err != nil {
		return nil, err
	}
	// 可能存在未知字段
	logrus.Info("upload response body ", res.String)

	var body UploadRes
	err = Unmarshal(res, &body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

// 订单流程状态
const (
	OrderStatusCreated  = 0  // 已创建
	OrderStatusProcess  = 3  // 处理中
	OrderStatusComplete = 4  // 处理成功
	OrderStatusFailed   = -1 // 处理失败
)

// GetResultRes 获取处理结果响应体
// 响应代码为成功时，才能确定订单状态
type GetResultRes struct {
	ResBody
	Content struct {
		OrderInfo struct {
			OrderId          string `json:"orderId"`          // 订单id
			FailType         int    `json:"failType"`         // 订单失败类型
			Status           int    `json:"status"`           // 订单流程状态
			OriginalDuration int    `json:"originalDuration"` // 上传设置音频时长，单位毫秒
			RealDuration     int    `json:"realDuration"`     // 实际处理音频时长，单位毫秒
		} `json:"orderInfo"` // 转写订单信息
		OrderResult      string `json:"orderResult"`      // 转写结果
		TaskEstimateTime int    `json:"taskEstimateTime"` // 订单预估耗时，单位毫秒
	} `json:"content"`
}

// GetResult 获取处理结果
func (a *xfyun) GetResult(orderId string) (*GetResultRes, error) {
	// 检查密钥设置
	err := a.Valid()
	if err != nil {
		return nil, err
	}

	// 查询字符串参数
	qs := a.SignA(a.lfAsrSecret)
	qs["orderId"] = orderId
	res, err := resty.Post(UrlGetResult, func(r *resty.Request) *resty.Request {
		r.SetHeader("Content-Type", "multipart/form-data")
		r.SetQueryParams(qs)
		return r
	})
	if err != nil {
		return nil, err
	}
	// 可能存在未知字段
	logrus.Info("get result response ", res.String)

	var body GetResultRes
	err = Unmarshal(res, &body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}
