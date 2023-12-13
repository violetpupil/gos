package logrus

import "github.com/rifflock/lfshook"

var (
	// 创建日志钩子
	NewHook = lfshook.NewHook
)
