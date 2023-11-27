package logrus

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var (
	// 创建日志器
	New = logrus.New

	// 默认日志器
	// 使用 logrus.New() 创建
	// Out os.Stderr
	// Formatter new(logrus.TextFormatter)
	// ReportCaller false
	//
	// 设置格式器
	SetFormatter = logrus.SetFormatter
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
	// Out 输出 writer
	// Formatter 格式器
	// ReportCaller 是否记录调用信息
	Logger = logrus.Logger
)

// Init 最佳选项设置
func Init() {
	// 包含调用信息
	logrus.SetReportCaller(true)
	DisableQuote()
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
