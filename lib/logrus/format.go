package logrus

import "github.com/sirupsen/logrus"

type (
	// 等号形式 空日志不显示 msg 字段
	// time="2023-11-27T14:17:08+08:00" level=info msg=hi
	// time="2023-11-27T14:21:00+08:00" level=info msg="{\"key\": 0}"
	//
	// DisableQuote 禁用引号
	// time=2023-11-27T14:22:39+08:00 level=info msg={"key": 0}
	TextFormatter = logrus.TextFormatter
	// json形式
	// {"level":"info","msg":"hi","time":"2023-11-27T14:30:04+08:00"}
	JSONFormatter = logrus.JSONFormatter
)

// DisableQuote 值不添加引号
func DisableQuote() {
	logrus.SetFormatter(
		&logrus.TextFormatter{
			DisableQuote: true,
		},
	)
}
