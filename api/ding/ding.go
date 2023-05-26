// 钉钉
// https://open.dingtalk.com/document/orgapp/learning-map
// https://open.dingtalk.com/document/orgapp/robot-overview
package ding

import (
	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/lib/resty"
)

const SendURL = "https://oapi.dingtalk.com/robot/send"

// Send 发送群消息
// 每个机器人每分钟最多发送20条消息到群里，如果超过20条，会限流10分钟。
// https://open.dingtalk.com/document/orgapp/custom-robots-send-group-messages
func Send(token string, body any) {
	res, err := resty.Post(SendURL, func(r *resty.Request) {
		r.SetQueryParam("access_token", token)
		r.SetBody(body)
	})
	if err != nil {
		logrus.Errorln("send ding error", err)
	} else {
		logrus.Infoln("send ding response", res.String())
	}
}

// SendClient 发送群消息，新创建resty客户端
// 每个机器人每分钟最多发送20条消息到群里，如果超过20条，会限流10分钟。
// https://open.dingtalk.com/document/orgapp/custom-robots-send-group-messages
func SendClient(token string, body any) {
	req := resty.New().R()
	req.SetQueryParam("access_token", token)
	req.SetBody(body)
	res, err := req.Post(SendURL)
	if err != nil {
		logrus.Errorln("send ding error", err)
	} else {
		logrus.Infoln("send ding response", res.String())
	}
}

// SendClientText 发送text类型群消息，新创建resty客户端
// 每个机器人每分钟最多发送20条消息到群里，如果超过20条，会限流10分钟。
// https://open.dingtalk.com/document/orgapp/custom-robots-send-group-messages
func SendClientText(token string, content string, atMobiles []string) {
	var body TextBody
	body.MsgType = "text"
	body.At.AtMobiles = atMobiles
	body.Text.Content = content

	req := resty.New().R()
	req.SetQueryParam("access_token", token)
	req.SetBody(body)
	res, err := req.Post(SendURL)
	if err != nil {
		logrus.Errorln("send ding error", err)
	} else {
		logrus.Infoln("send ding response", res.String())
	}
}
