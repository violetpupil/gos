package logrus

import "github.com/sirupsen/logrus"

type (
	// 文本形式
	// time="2023-11-27T14:17:08+08:00" level=info msg=hi
	// 空日志不显示 msg 字段
	// 有时会添加引号 msg="{\"key\": 0}"
	// 检测到 tty 时，会输出颜色，记录程序运行秒数 INFO[0000] hi
	//
	// ForceColors 强制以 tty 格式输出
	// FullTimestamp 记录时间戳
	// TimestampFormat 时间戳格式
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
