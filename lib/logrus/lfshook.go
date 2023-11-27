package logrus

import "github.com/rifflock/lfshook"

var (
	// 创建分级日志钩子
	NewHook = lfshook.NewHook
)
