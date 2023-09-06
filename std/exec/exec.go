package exec

import (
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

// RunCommand 运行命令直到完成，并获取输出
func RunCommand(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	var out strings.Builder
	var e strings.Builder
	cmd.Stdout = &out
	cmd.Stderr = &e

	err := cmd.Run()
	// 打印执行时的错误
	if e.Len() > 0 {
		logrus.Errorln("cmd error", e.String())
	}
	if err != nil {
		logrus.Errorln("cmd run error", err)
		return "", err
	}
	return out.String(), nil
}
