// 钉钉
// https://open.dingtalk.com/document/orgapp/learning-map
// https://open.dingtalk.com/document/orgapp/robot-overview
package ding

const SendURL = "https://oapi.dingtalk.com/robot/send"

// Send 发送群消息
// 每个机器人每分钟最多发送20条消息到群里，如果超过20条，会限流10分钟。
// https://open.dingtalk.com/document/orgapp/custom-robots-send-group-messages
func Send(token string, text string) error {
	return nil
}
