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
