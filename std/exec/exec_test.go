package exec

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestRunCommand(t *testing.T) {
	// 防止打印标准错误时，因引号导致换行不生效
	logrus.SetFormatter(
		// 等号形式 - 默认
		&logrus.TextFormatter{
			DisableQuote: true,
		},
	)
	r, err := RunCommand("ipconfig")
	fmt.Println(r, err)
}
