package ocr

import (
	"fmt"
	"time"

	"github.com/violetpupil/gos/lib/resty"
	"go.uber.org/zap"
)

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

// 访问ocr接口用的token
// https://ai.baidu.com/ai-doc/REFERENCE/Ck3dwjhhu
func BaiduToken(trace, apiKey, secretKey string) (string, error) {
	log := zap.L().With(zap.String("traceId", trace))

	var suss TokenResult
	var fail TokenError
	res, err := resty.New().
		SetTimeout(5 * time.Second).
		R().
		SetQueryParams(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     apiKey,
			"client_secret": secretKey,
		}).
		SetResult(&suss).
		SetError(&fail).
		Post("https://aip.baidubce.com/oauth/2.0/token")
	if err != nil {
		log.Error("post token error", zap.Error(err))
		return "", err
	}

	if res.IsSuccess() {
		log.Info("post token success")
		return suss.AccessToken, nil
	} else {
		log.Error("post token fail", zap.Any("body", fail))
		return "", fmt.Errorf("%s: %s", fail.Error, fail.ErrorDescription)
	}
}
