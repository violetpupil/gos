package logrus

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var (
	// 两个操作数都不是字符串时，会添加空格
	Info  = logrus.Info
	Infof = logrus.Infof
	// 总是在两个操作数之间添加空格
	Infoln  = logrus.Infoln
	Error   = logrus.Error
	Errorln = logrus.Errorln
	// Fatal 记录Fatal级别日志，然后退出程序
	Fatal   = logrus.Fatal
	Fatalln = logrus.Fatalln
)

type (
	// 日志
	Entry = logrus.Entry
	// 日志器
	Logger = logrus.Logger
)

// Init 最佳选项设置
func Init() {
	// 包含调用信息
	logrus.SetReportCaller(true)
	DisableQuote()
}

// DisableQuote 值不添加引号
func DisableQuote() {
	// 日志格式，空日志不显示msg字段
	// time=2023-03-10T15:43:04+08:00 level=info msg=log
	logrus.SetFormatter(
		// 等号形式 - 默认
		&logrus.TextFormatter{
			DisableQuote: true,
		},
	)
}

// Type 打印接口实际类型
func Type(value any) {
	t := fmt.Sprintf("%T", value)
	logrus.WithField("Type", t).Info("check type ", value)
}

// WithStruct 添加结构体字段，打印键名
func WithStruct(key string, value any) *logrus.Entry {
	return logrus.WithField(key, fmt.Sprintf("%+v", value))
}
