package smtp

import (
	"fmt"
	"net/smtp"

	"github.com/violetpupil/components/mapping"
)

// SendMail 发送邮件
// 参考阿里云 smtp 调用示例
// https://help.aliyun.com/document_detail/29457.html
func SendMail(
	username, password, host, port, to, from, subject, body string,
) error {
	headers := map[string]string{
		"To":           to,
		"From":         from,
		"Subject":      subject,
		"Content-Type": "text/html; charset=UTF-8",
	}
	header := mapping.Join(headers, ": ", "\r\n")
	msg := fmt.Sprint(header, "\r\n\r\n", body)

	addr := fmt.Sprint(host, ":", port)
	auth := smtp.PlainAuth("", username, password, host)
	err := smtp.SendMail(addr, auth, username, []string{to}, []byte(msg))
	return err
}
