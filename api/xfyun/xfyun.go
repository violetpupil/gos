// 科大讯飞api客户端
// 如果不使用 NewXfyun，必须先 Init 初始化
// 不同服务的密钥不同
// https://www.xfyun.cn/doc/
// https://www.xfyun.cn/document/error-code 错误码查询
package xfyun

import (
	"encoding/json"

	"github.com/violetpupil/components/api/xfyun/asr"
	"github.com/violetpupil/components/api/xfyun/trans"
	"github.com/violetpupil/components/lib/logrus"
)

type Client struct {
	// 配置
	Appid       string `json:"appid"`
	ApiSecret   string `json:"apiSecret"`   // api密钥值
	ApiKey      string `json:"apiKey"`      // api密钥键
	LfAsrSecret string `json:"lfAsrSecret"` // 语音转写密钥

	// 服务集成
	LfAsr    asr.LfAsr      `json:"lfAsr"`    // 语音转写
	NiuTrans trans.NiuTrans `json:"niuTrans"` // 机器翻译
}

// NewXfyun 创建科大讯飞api客户端，cfg是配置json字符串
func NewXfyun(cfg string) (*Client, error) {
	c := new(Client)
	err := json.Unmarshal([]byte(cfg), c)
	if err != nil {
		logrus.Error("json unmarshal error ", err)
		return nil, err
	}

	c.LfAsr.Appid = c.Appid
	c.LfAsr.LfAsrSecret = c.LfAsrSecret
	c.NiuTrans.Appid = c.Appid
	c.NiuTrans.ApiKey = c.ApiKey
	c.NiuTrans.ApiSecret = c.ApiSecret
	return c, nil
}

var Xfyun *Client

// Init 初始化api客户端
func Init(appid string) {
	Xfyun = &Client{
		Appid:    appid,
		LfAsr:    asr.LfAsr{Appid: appid},
		NiuTrans: trans.NiuTrans{Appid: appid},
	}
}

// SetLfAsrSecret 设置语音转写密钥
func SetLfAsrSecret(s string) {
	Xfyun.LfAsr.LfAsrSecret = s
}

// SetApiSecret 设置api密钥
func SetApiSecret(k, s string) {
	Xfyun.NiuTrans.ApiKey = k
	Xfyun.NiuTrans.ApiSecret = s
}
