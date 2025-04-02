package ocr

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/violetpupil/gos/lib/resty"
	"go.uber.org/zap"
)

type BaiduToken struct {
	APIKey            string
	SecretKey         string
	AccessToken       string
	AccessTokenExpire int64 // access token过期时间戳
}

func NewBaiduToken(apiKey, secretKey string) *BaiduToken {
	return &BaiduToken{
		APIKey:    apiKey,
		SecretKey: secretKey,
	}
}

// TokenResult 获取access token成功响应
type BaiduTokenResult struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"` // 多少秒后过期
}

// TokenError 获取access token失败响应
type BaiduTokenError struct {
	Error            string `json:"error"`             // 错误类型
	ErrorDescription string `json:"error_description"` // 错误描述
}

// 访问ocr接口用的token
// https://ai.baidu.com/ai-doc/REFERENCE/Ck3dwjhhu
func (b *BaiduToken) Get(trace string) (string, error) {
	log := zap.L().With(zap.String("traceId", trace))

	// 检查是否过期
	if b.AccessToken != "" && time.Now().Unix() < b.AccessTokenExpire {
		log.Info("access token valid")
		return b.AccessToken, nil
	}

	// 重新申请
	var suss BaiduTokenResult
	var fail BaiduTokenError
	res, err := resty.New().
		SetTimeout(5 * time.Second).
		R().
		SetQueryParams(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     b.APIKey,
			"client_secret": b.SecretKey,
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
		b.AccessToken = suss.AccessToken
		b.AccessTokenExpire = time.Now().Unix() + suss.ExpiresIn
		return suss.AccessToken, nil
	} else {
		log.Error("post token fail", zap.Any("body", fail))
		return "", fmt.Errorf("%s: %s", fail.Error, fail.ErrorDescription)
	}
}

const (
	BaiduOCRModeAccurate = "accurate" // 通用文字识别（高精度含位置版）
	BaiduOCRModeGeneral  = "general"  // 通用文字识别（标准含位置版）
)

type OCRResult struct {
	LogID          uint64        `json:"log_id"`
	WordsResultNum uint32        `json:"words_result_num"`
	WordsResult    []WordsResult `json:"words_result"`
	ErrorCode      int           `json:"error_code"` // 错误码
	ErrorMsg       string        `json:"error_msg"`  // 错误描述
}

type WordsResult struct {
	Words    string   `json:"words"`
	Location Location `json:"location"`
}

type Location struct {
	Top    uint32 `json:"top"`
	Left   uint32 `json:"left"`
	Width  uint32 `json:"width"`
	Height uint32 `json:"height"`
}

// OCR 通用文字识别
// 通用文字识别（高精度含位置版） https://ai.baidu.com/ai-doc/OCR/tk3h7y2aq
// 通用文字识别（标准含位置版） https://ai.baidu.com/ai-doc/OCR/vk3h7y58v
func BaiduOCR(trace, apiKey, secretKey, mode string, image string) (*OCRResult, error) {
	log := zap.L().With(zap.String("traceId", trace))

	token, err := NewBaiduToken(apiKey, secretKey).Get(trace)
	if err != nil {
		return nil, err
	}

	var result OCRResult
	res, err := resty.New().
		SetTimeout(10*time.Second).
		R().
		SetQueryParam("access_token", token).
		SetFormData(map[string]string{"image": image}).
		SetResult(&result).
		Post("https://aip.baidubce.com/rest/2.0/ocr/v1/" + mode)
	if err != nil {
		log.Error("post ocr error", zap.Error(err))
		return nil, err
	}

	// 错误的响应码也是200
	if res.StatusCode() != http.StatusOK {
		log.Error("post ocr status fail", zap.String("status", res.Status()))
		return nil, errors.New(res.Status())
	} else if result.ErrorCode == 0 {
		log.Info("post ocr response", zap.String("body", res.String()))
		return &result, nil
	} else {
		log.Error("post ocr fail", zap.String("body", res.String()))
		return nil, fmt.Errorf("%d: %s", result.ErrorCode, result.ErrorMsg)
	}
}
