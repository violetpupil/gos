// https://ai.baidu.com/ai-doc/index/OCR
package baidu

import (
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type OCRResult struct {
	LogID          uint64        `json:"log_id"`
	WordsResultNum uint32        `json:"words_result_num"`
	WordsResult    []WordsResult `json:"words_result"`
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
func (b *Baidu) OCR(image string) (*OCRResult, error) {
	var result OCRResult
	res, err := resty.New().R().
		SetQueryParam("access_token", "").
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
	return &result, nil
}
