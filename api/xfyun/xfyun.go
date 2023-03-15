// 科大讯飞api客户端
// 如果不使用 NewXfyun，必须先 Init 初始化
// 不同服务的密钥不同
// https://www.xfyun.cn/doc/
// https://www.xfyun.cn/document/error-code 错误码查询
package xfyun

import (
	"encoding/json"
	"os"
)

type Client struct {
	Appid    string   `json:"appid"`
	LfAsr    lfAsr    `json:"lfAsr"`    // 语音转写
	NiuTrans niuTrans `json:"niuTrans"` // 机器翻译
}

// NewXfyun 创建科大讯飞api客户端，cfg是配置json字符串
func NewXfyun(cfg string) (*Client, error) {
	c := new(Client)
	err := json.Unmarshal([]byte(cfg), c)
	return c, err
}

var Xfyun *Client

// Init 初始化api客户端
func Init(appid string) {
	Xfyun = &Client{
		Appid:    appid,
		LfAsr:    lfAsr{Appid: appid},
		NiuTrans: niuTrans{Appid: appid},
	}
}

// InitEnv 用环境变量初始化api客户端
func InitEnv() {
	Init(os.Getenv("XfyunAppid"))
	SetLfAsrSecret(os.Getenv("XfyunLfAsrSecret"))
	SetTransSecret(os.Getenv("XfyunTransKey"), os.Getenv("XfyunTransSecret"))
}

// SetLfAsrSecret 设置语音转写密钥
func SetLfAsrSecret(s string) {
	Xfyun.LfAsr.LfAsrSecret = s
}

// SetTransSecret 设置机器翻译密钥
func SetTransSecret(k, s string) {
	Xfyun.NiuTrans.TransKey = k
	Xfyun.NiuTrans.TransSecret = s
}
