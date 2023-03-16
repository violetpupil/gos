// 机器翻译
// 如果不使用 NewXfyun，必须先 SetTransSecret 设置密钥
// https://www.xfyun.cn/doc/nlp/niutrans/API.html
package trans

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/lib/resty"
	"github.com/violetpupil/components/std/hmac"
	"github.com/violetpupil/components/std/sha256"
)

// 语种
const (
	LanguageCn  = "cn"  // 中文(简体)
	LanguageCht = "cht" // 中文(繁体)
	LanguageEn  = "en"  // 英语
)

type NiuTrans struct {
	Appid       string `json:"appid" env:"XfyunAppid"`
	TransSecret string `json:"transSecret" env:"XfyunTransSecret"` // 机器翻译密钥值
	TransKey    string `json:"transKey" env:"XfyunTransKey"`       // 机器翻译密钥键
}

// NewNiuTrans 创建科大讯飞机器翻译客户端，cfg是配置json字符串
func NewNiuTrans(cfg string) (*NiuTrans, error) {
	a := new(NiuTrans)
	err := json.Unmarshal([]byte(cfg), a)
	return a, err
}

const (
	UrlTrans = "https://ntrans.xfyun.cn/v2/ots"
	// 签名请求头主机
	SignHost = "ntrans.xfyun.cn"
	// 签名请求行
	SignLine = "POST /v2/ots HTTP/1.1"
	// 单次翻译最大字符数，包括符号
	MaxChars = 5000
)

// TranslateReq 机器翻译请求体
type TranslateReq struct {
	Common struct {
		Appid string `json:"app_id"`
	} `json:"common"`
	Business struct {
		From string `json:"from"` // 源语种 可以指定auto自动识别源语种
		To   string `json:"to"`   // 目标语种
	} `json:"business"`
	Data struct {
		// 源文本 utf8字符 base64编码
		// 编码后字节数不超过 20000
		Text string `json:"text"`
	} `json:"data"`
}

// TranslateRes 机器翻译响应体
// 响应码不是200时，只有message字段
// 返回码异常时，没有data字段
type TranslateRes struct {
	Code    int    `json:"code"`    // 返回码，0表示成功，其它表示异常
	Message string `json:"message"` // 描述信息
	Sid     string `json:"sid"`     // 本次会话id
	Data    struct {
		Result struct {
			From        string `json:"from"` // 源语种，如果请求设置auto则自动返回识别到的源语种参数
			To          string `json:"to"`   // 目标语种
			TransResult struct {
				Dst string `json:"dst"` // 翻译文本
				Src string `json:"src"` // 源文本
			} `json:"trans_result"`
		} `json:"result"`
	} `json:"data"`
}

// StatusError 响应码不是200时，获取错误
func (res TranslateRes) StatusError(status string) error {
	return fmt.Errorf("translate status: %s, message: %s", status, res.Message)
}

// Error 返回码异常时，获取错误
func (res TranslateRes) Error() error {
	e := fmt.Errorf("translate code: %d, message: %s, sid: %s", res.Code, res.Message, res.Sid)
	return e
}

// Translate 机器翻译 src源语种 dst目标语种 返回翻译文本
func (a NiuTrans) Translate(text, src, dst string) (string, error) {
	// 检查密钥设置
	if a.Appid == "" || a.TransSecret == "" || a.TransKey == "" {
		return "", fmt.Errorf("argument insufficient %+v", a)
	}

	// 请求体
	var bodyS TranslateReq
	bodyS.Common.Appid = a.Appid
	bodyS.Business.From = src
	bodyS.Business.To = dst
	bodyS.Data.Text = base64.StdEncoding.EncodeToString([]byte(text))
	body, err := json.Marshal(bodyS)
	if err != nil {
		logrus.Error("json marshal error ", err)
		return "", err
	}

	headers := a.Sign(SignHost, SignLine, body)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json,version=1.0"
	res, err := resty.Post(UrlTrans, func(r *resty.Request) *resty.Request {
		r.SetHeaders(headers)
		r.SetBody(body)
		return r
	})
	if err != nil {
		logrus.Error("post error ", err)
		return "", err
	}
	// 打印所有字段
	logrus.Info("translate response body ", res.String)

	var resBody TranslateRes
	err = json.Unmarshal(res.Body, &resBody)
	if err != nil {
		logrus.Error("json unmarshal error ", err)
		return "", res.ToError()
	}
	// 请求失败
	if res.StatusCode != 200 {
		return "", resBody.StatusError(res.Status)
	}
	// 处理失败
	if resBody.Code != 0 {
		return "", resBody.Error()
	}
	result := resBody.Data.Result.TransResult.Dst
	return result, nil
}

// Sign 生成签名请求头
// host 为请求头主机，line为请求行，body为请求体
func (a NiuTrans) Sign(host, line string, body []byte) map[string]string {
	// 服务器时间不准可能导致请求失败
	date := time.Now().UTC().Format(time.RFC1123)
	// 请求体摘要
	digest := sha256.Sum256Base64(body)
	digest = "SHA-256=" + digest
	// 签名host date digest请求头 + 请求行
	parts := strings.Join([]string{
		"host: " + host,
		"date: " + date,
		line,
		"digest: " + digest,
	}, "\n")
	sign := hmac.HmacSha256Base64(a.TransSecret, parts)
	authorization := strings.Join([]string{
		`api_key="%s"`,
		`algorithm="hmac-sha256"`,
		`headers="host date request-line digest"`,
		`signature="%s"`,
	}, ", ")
	authorization = fmt.Sprintf(authorization, a.TransKey, sign)

	headers := map[string]string{
		"Host":          host,
		"Date":          date,
		"Digest":        digest,
		"Authorization": authorization,
	}
	return headers
}
