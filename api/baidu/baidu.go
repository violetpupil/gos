// 百度客户端
// https://ai.baidu.com/ai-doc
package baidu

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v7"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type Client struct {
	APIKey            string `env:"BAIDU_API_KEY"`
	SecretKey         string `env:"BAIDU_SECRET_KEY"`
	AccessToken       string
	AccessTokenExpire int64 // access token过期时间戳
}

// 百度客户端，必须先调用Init或InitEnv初始化
var Baidu *Client

// Init 初始化百度客户端
func Init(api, secret string) *Client {
	Baidu = new(Client)
	Baidu.APIKey = api
	Baidu.SecretKey = secret
	return Baidu
}

// InitEnv 初始化百度客户端 使用环境变量
func InitEnv() error {
	Baidu = new(Client)
	err := env.Parse(Baidu)
	return err
}

// TokenResult 获取access token成功响应
type TokenResult struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"` // 多少秒后过期
}

// TokenError 获取access token失败响应
type TokenError struct {
	Error            string `json:"error"`             // 错误类型
	ErrorDescription string `json:"error_description"` // 错误描述
}

// GetAccessToken 获取access token
// https://ai.baidu.com/ai-doc/REFERENCE/Ck3dwjhhu
func (b *Client) GetAccessToken() (string, error) {
	// 检查是否过期
	if b.AccessToken != "" && time.Now().Unix() < b.AccessTokenExpire {
		logrus.Infoln("access token valid")
		return b.AccessToken, nil
	}

	// 重新申请
	var suss TokenResult
	var fail TokenError
	res, err := resty.New().R().
		SetQueryParams(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     Baidu.APIKey,
			"client_secret": Baidu.SecretKey,
		}).SetResult(&suss).
		SetError(&fail).
		Post("https://aip.baidubce.com/oauth/2.0/token")
	if err != nil {
		logrus.Errorln("post token error", err)
		return "", err
	}
	logrus.WithFields(logrus.Fields{
		"status": res.Status(),
		"body":   res.String(),
	}).Infoln("post token response")

	if res.IsSuccess() {
		b.AccessToken = suss.AccessToken
		b.AccessTokenExpire = time.Now().Unix() + suss.ExpiresIn
		logrus.Infoln("post token success", suss.AccessToken)
		return b.AccessToken, nil
	} else {
		err = fmt.Errorf("%s: %s", fail.Error, fail.ErrorDescription)
		logrus.Errorln("post token fail", err)
		return "", err
	}
}
