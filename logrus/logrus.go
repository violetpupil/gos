package logrus

import "github.com/sirupsen/logrus"

func Init() {
	// 包含调用信息
	logrus.SetReportCaller(true)
}
