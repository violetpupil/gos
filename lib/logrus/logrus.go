package logrus

import "github.com/sirupsen/logrus"

var (
	Info  = logrus.Info
	Infof = logrus.Infof
	Error = logrus.Error
	// Fatal 记录Fatal级别日志，然后退出程序
	Fatal = logrus.Fatal
)

func Init() {
	// 包含调用信息
	logrus.SetReportCaller(true)

	// 日志格式，空日志不显示msg字段
	// time=2023-03-10T15:43:04+08:00 level=info msg=log
	logrus.SetFormatter(
		// 等号形式 - 默认
		&logrus.TextFormatter{
			// 值不添加引号
			DisableQuote: true,
		},
	)
}
