package ding

// Body 基本请求体
type Body struct {
	MsgType string `json:"msgtype"` // 消息类型
}

// TextBody text类型请求体
type TextBody struct {
	Body
	At struct {
		AtMobiles []string `json:"atMobiles"` // 被@人的手机号
	} `json:"at"`
	Text struct {
		Content string `json:"content"` // 消息内容
	} `json:"text"`
}
