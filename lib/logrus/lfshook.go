package logrus

import "github.com/rifflock/lfshook"

var (
	// 创建本地文件系统日志钩子
	NewHook = lfshook.NewHook
)
