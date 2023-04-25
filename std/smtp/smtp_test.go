package smtp

import (
	"fmt"
	"testing"
)

func TestSendMail(t *testing.T) {
	err := SendMail(
		"",
		"",
		"",
		"",
		"495140158@qq.com",
		"hello",
		"<html><body>hi</body></html>",
	)
	fmt.Println(err)
}
