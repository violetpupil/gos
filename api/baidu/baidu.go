// 百度客户端，必须先调用Init初始化
// https://ai.baidu.com/ai-doc
package baidu

import (
	"fmt"

	"github.com/caarlos0/env/v7"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type baidu struct {
	APIKey    string `env:"BAIDU_API_KEY"`
	SecretKey string `env:"BAIDU_SECRET_KEY"`
}

var Baidu = new(baidu)

// Init 初始化百度客户端
func Init() error {
	err := env.Parse(Baidu)
	return err
}

// GenAccessToken 获取access token
// https://ai.baidu.com/ai-doc/REFERENCE/Ck3dwjhhu
func GenAccessToken() error {
	res, err := resty.New().R().
		SetQueryParams(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     Baidu.APIKey,
			"client_secret": Baidu.SecretKey,
		}).Post("https://aip.baidubce.com/oauth/2.0/token")
	if err != nil {
		logrus.Errorln("post token error", err)
		return err
	}

	fmt.Println(res.String())
	return nil
}
