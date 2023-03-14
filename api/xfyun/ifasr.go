// 语音转写 Long Form Automatic Speech Recognition
// 必须先 SetLfAsrSecret 设置语音转写密钥
// 上传音视频文件，可以设置处理结果回调，再调用结果查询接口
// https://www.xfyun.cn/doc/asr/ifasr_new/API.html
package xfyun

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/lib/resty"
	"github.com/violetpupil/components/media/srt"
	"github.com/violetpupil/components/std/json"
	"github.com/violetpupil/components/std/os"
	"github.com/violetpupil/components/std/strconv"
	"github.com/violetpupil/components/std/time"
)

const (
	// 文件上传地址
	UrlUpload = "https://raasr.xfyun.cn/v2/api/upload"
	// 获取结果地址
	UrlGetResult = "https://raasr.xfyun.cn/v2/api/getResult"
)

// Valid 验证语音转写设置是否有效
func (a *xfyun) Valid() error {
	if a == nil || a.appid == "" || a.lfAsrSecret == "" {
		return fmt.Errorf("argument insufficient %+v", a)
	}
	return nil
}

// SpeechToText 语音转写，返回 srt 字幕
// name 是本地音视频文件路径，timeout 是轮询获取识别结果超时秒数
func (a *xfyun) SpeechToText(name string, timeout int) (string, error) {
	return "", nil
}

// UploadRes 上传音视频文件响应体
type UploadRes struct {
	ResBody
	Content struct {
		OrderId          string `json:"orderId"`          // 订单id
		TaskEstimateTime int    `json:"taskEstimateTime"` // 订单预估耗时，单位毫秒
	} `json:"content"`
}

// Upload 上传音视频文件
func (a *xfyun) Upload(name string) (*UploadRes, error) {
	// 检查密钥设置
	err := a.Valid()
	if err != nil {
		logrus.Error("valid error ", err)
		return nil, err
	}

	// 查询字符串参数
	// callbackUrl 订单完成时通知回调地址
	info, err := os.Stat(name)
	if err != nil {
		logrus.Error("stat error ", err)
		return nil, err
	}
	qs := a.SignA(a.lfAsrSecret)
	qs["fileName"] = info.Name
	qs["fileSize"] = strconv.ToString(info.Size)
	// 音频时长，单位为毫秒
	// 必须有值，设置为0订单也可以成功
	qs["duration"] = "0"

	// 请求体
	bytes, err := os.ReadFile(name)
	if err != nil {
		logrus.Error("read file error ", err)
		return nil, err
	}
	res, err := resty.Post(UrlUpload, func(r *resty.Request) *resty.Request {
		r.SetHeader("Content-Type", "application/json")
		r.SetQueryParams(qs)
		r.SetBody(bytes)
		return r
	})
	if err != nil {
		logrus.Error("post error ", err)
		return nil, err
	}
	// 打印所有字段
	logrus.Info("upload response body ", res.String)

	var body UploadRes
	err = Unmarshal(res, &body)
	if err != nil {
		logrus.Error("unmarshal error ", err)
		return nil, err
	}
	return &body, nil
}

// UploadCallback 订单完成时通知回调处理 get请求
func (a *xfyun) UploadCallback(c *gin.Context) {
	// qs参数
	logrus.Info("upload callback url ", c.Request.URL)
	// 订单id
	orderId := c.Query("orderId")
	// 订单状态 1 转写成功 -1 转写失败
	status := c.Query("status")
	if orderId == "" {
		logrus.Info("callback order id empty")
		return
	}
	if status == "-1" {
		logrus.Info("recognition failed")
		return
	}
	if status != "1" {
		logrus.Info("unknown order status")
		return
	}
	logrus.Info("recognition success")

	// 查询结果并打印句子
	res, err := a.GetResult(orderId)
	if err != nil {
		logrus.Error("get result error ", err)
		return
	}
	err = a.OrderResult(res)
	logrus.Info("order result error ", err)
}

// 订单流程状态
const (
	OrderStatusCreated  = 0  // 已创建
	OrderStatusProcess  = 3  // 处理中
	OrderStatusComplete = 4  // 处理完成-成功
	OrderStatusFailed   = -1 // 处理完成-失败
)

// GetResultRes 获取处理结果响应体
// 响应代码为成功时，才能确定订单状态
type GetResultRes struct {
	ResBody
	Content struct {
		OrderInfo struct {
			OrderId          string `json:"orderId"`          // 订单id
			FailType         int    `json:"failType"`         // 订单失败类型，处理成功时响应0
			Status           int    `json:"status"`           // 订单流程状态
			OriginalDuration int    `json:"originalDuration"` // 上传设置音频时长，单位毫秒
			RealDuration     int    `json:"realDuration"`     // 实际处理音频时长，单位毫秒
			ExpireTime       int    `json:"expireTime"`       // 已完成订单删除时间戳，毫秒
		} `json:"orderInfo"` // 转写订单信息
		OrderResult      string `json:"orderResult"`      // 转写结果，json字符串
		TaskEstimateTime int    `json:"taskEstimateTime"` // 订单预估剩余耗时，单位毫秒
	} `json:"content"`
}

// OrderResult 转写结果
// lattice 中 json_1best 是字符串
// lattice2 中 json_1best 是对象
type OrderResult struct {
	Lattice  []Lattice  `json:"lattice"`  // 顺滑处理后的结果
	Lattice2 []Lattice2 `json:"lattice2"` // 未顺滑处理的结果
}

// Lattice 单句转写结果
type Lattice struct {
	Json1best Json1best `json:"json_1best"`
}

// UnmarshalJSON json解码
// json_1best 响应字段是json字符串，将其直接解码为对象
func (l *Lattice) UnmarshalJSON(data []byte) error {
	type tmp struct {
		Json1best string `json:"json_1best"`
	}

	var t tmp
	err := json.Unmarshal(data, &t)
	if err != nil {
		logrus.Error("json unmarshal error ", err)
		return err
	}
	err = json.Unmarshal([]byte(t.Json1best), &l.Json1best)
	return err
}

// Lattice2 单句转写结果
type Lattice2 struct {
	Json1best Json1best `json:"json_1best"`
}

// Json1best 单句转写结果
type Json1best struct {
	St St `json:"st"`
}

// St 单句转写结果
type St struct {
	Bg string `json:"bg"` // 单句开始时间毫秒数
	Ed string `json:"ed"` // 单句结束时间毫秒数
	Rl string `json:"rl"`
	Rt []Rt   `json:"rt"` // 词语识别结果
}

// Rt 词语识别结果
type Rt struct {
	Ws []Ws `json:"ws"`
}

// Ws 词语识别结果
type Ws struct {
	Wb int64 `json:"wb"` // 词语相对St.Bg开始帧数 一帧10ms
	We int64 `json:"we"` // 词语相对St.Bg结束帧数 一帧10ms
	Cw []Cw  `json:"cw"` // 词语识别结果
}

// Cw 词语识别结果
type Cw struct {
	W  string `json:"w"`  // 识别结果
	Wp string `json:"wp"` // 词语类型
}

// 词语类型
const (
	WpNormal = "n" // 正常词
	WpSmooth = "s" // 顺滑
	WpPunct  = "p" // 标点
	WpSeg    = "g" // 分段标识，对应Cw.W为空
)

// GetResult 获取处理结果，处理完成后72小时可查
func (a *xfyun) GetResult(orderId string) (*GetResultRes, error) {
	// 检查密钥设置
	err := a.Valid()
	if err != nil {
		logrus.Error("valid error ", err)
		return nil, err
	}

	// 查询字符串参数
	qs := a.SignA(a.lfAsrSecret)
	qs["orderId"] = orderId
	res, err := resty.Post(UrlGetResult, func(r *resty.Request) *resty.Request {
		r.SetHeader("Content-Type", "application/json")
		r.SetQueryParams(qs)
		return r
	})
	if err != nil {
		logrus.Error("post error ", err)
		return nil, err
	}
	// 打印所有字段
	logrus.Info("get result response ", res.String)

	var body GetResultRes
	err = Unmarshal(res, &body)
	if err != nil {
		logrus.Error("unmarshal error ", err)
		return nil, err
	}
	return &body, nil
}

// OrderResult 解码转写结果
func (a *xfyun) OrderResult(res *GetResultRes) error {
	var result OrderResult
	err := json.Unmarshal([]byte(res.Content.OrderResult), &result)
	if err != nil {
		logrus.Error("json unmarshal error ", err)
		return err
	}
	err = a.WriteSrt(res.Content.OrderInfo.OrderId, result.Lattice)
	return err
}

// Sentence 拼接句子输出
func (a *xfyun) Sentence(las []Lattice) {
	for _, la := range las {
		for _, rt := range la.Json1best.St.Rt {
			for _, ws := range rt.Ws {
				if len(ws.Cw) > 0 {
					cw := ws.Cw[0]
					if cw.Wp == WpSeg {
						fmt.Println()
					} else {
						fmt.Print(cw.W)
					}
				}
			}
		}
	}
}

// WriteSrt 写srt字幕文件 订单id.srt
func (a *xfyun) WriteSrt(orderId string, las []Lattice) error {
	subs := make([]srt.Subtitle, 0)
	// 处理每句
	for i, la := range las {
		var line string
		// 处理每个词
		for _, rt := range la.Json1best.St.Rt {
			for _, ws := range rt.Ws {
				if len(ws.Cw) > 0 {
					cw := ws.Cw[0]
					switch cw.Wp {
					case WpNormal, WpSmooth:
						line += cw.W
					// 标点用空格代替
					case WpPunct:
						line += " "
					}
				}
			}
		}

		if line != "" {
			// 将毫秒数转为时间字符串
			bg, err := a.VideoTime(la.Json1best.St.Bg)
			if err != nil {
				// 会导致字幕缺失
				logrus.Error("bg video time error ", err)
				continue
			}
			ed, err := a.VideoTime(la.Json1best.St.Ed)
			if err != nil {
				// 会导致字幕缺失
				logrus.Error("ed video time error ", err)
				continue
			}
			sub := srt.Subtitle{
				Id:    i + 1,
				Start: bg,
				End:   ed,
				Line:  []string{line},
			}
			subs = append(subs, sub)
		}
	}

	name := orderId + ".srt"
	err := srt.WriteSrt(subs, name)
	return err
}

// VideoTime 将毫秒数转为视频时间字符串
func (a *xfyun) VideoTime(msec string) (string, error) {
	ms, err := strconv.ParseInt(msec, 10, 64)
	if err != nil {
		logrus.Error("parse int error ", err)
		return "", err
	}
	t := time.Video(ms)
	return t, nil
}
