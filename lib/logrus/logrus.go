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
}
