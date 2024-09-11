package utils

import (
	"log"

	"go.uber.org/zap"
)

func InitDevelopmentLog() {
	c := zap.NewDevelopmentConfig()
	// 不输出调用栈
	c.DisableStacktrace = true

	l, err := c.Build()
	if err != nil {
		log.Fatalln("new development error", err)
	} else {
		zap.ReplaceGlobals(l)
	}
}

// SyncLog 刷写全局logger缓存的日志，程序结束前调用
func SyncLog() {
	err := zap.L().Sync()
	if err != nil {
		log.Println("sync log error", err)
	}
}
