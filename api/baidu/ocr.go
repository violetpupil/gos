// https://ai.baidu.com/ai-doc/index/OCR
package baidu

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
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
// https://ai.baidu.com/ai-doc/OCR/vk3h7y58v
func (b *client) OCR(image string) (*OCRResult, error) {
	token, err := b.GetAccessToken()
	if err != nil {
		logrus.Errorln("get access token error", err)
		return nil, err
	}

	var result OCRResult
	res, err := resty.New().R().
		SetQueryParam("access_token", token).
		SetFormData(map[string]string{"image": image}).
		SetResult(&result).
		Post("https://aip.baidubce.com/rest/2.0/ocr/v1/general")
	if err != nil {
		logrus.Errorln("post ocr error", err)
		return nil, err
	}
	logrus.WithFields(logrus.Fields{
		"status": res.Status(),
		"body":   res.String(),
	}).Infoln("post ocr response")

	if res.StatusCode() != http.StatusOK {
		logrus.Errorln("post ocr status fail", res.Status())
		return nil, errors.New(res.Status())
	}
	// 错误的响应码也是200
	if result.ErrorCode == 0 {
		logrus.Infof("post ocr success %+v", result)
		return &result, nil
	} else {
		err = fmt.Errorf("%d: %s", result.ErrorCode, result.ErrorMsg)
		logrus.Errorln("post ocr fail", err)
		return nil, err
	}
}
