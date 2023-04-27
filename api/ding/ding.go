// 钉钉
// https://open.dingtalk.com/document/orgapp/learning-map
// https://open.dingtalk.com/document/orgapp/robot-overview
package ding

import (
	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/lib/resty"
)

// Send 发送群消息
// 每个机器人每分钟最多发送20条消息到群里，如果超过20条，会限流10分钟。
// https://open.dingtalk.com/document/orgapp/custom-robots-send-group-messages
func Send(token string, body any) error {
	res, err := resty.Post(
		"https://oapi.dingtalk.com/robot/send",
		func(r *resty.Request) *resty.Request {
			r.SetQueryParam("access_token", token)
			r.SetBody(body)
			return r
		},
	)
	if err != nil {
		logrus.Errorln("post error", err)
		return err
	}

	logrus.WithFields(logrus.Fields{
		"Status": res.Status(),
		"Body":   res.String(),
	}).Info("send response")
	return nil
}
