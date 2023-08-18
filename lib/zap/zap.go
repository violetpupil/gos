// zap好像比logrus块
// zap必须要初始化
package zap

import (
	"log"

	"go.uber.org/zap"
)

var (
	// 全局logger
	L = zap.L

	// 字段
	Any      = zap.Any      // any接口
	String   = zap.String   // 字符串
	Int      = zap.Int      // 整型
	Duration = zap.Duration // 持续时间
	Error    = zap.Error    // 错误，键名error
)

// InitDevelopment 将zap全局logger设置为开发
func InitDevelopment() error {
	l, err := zap.NewDevelopment()
	if err != nil {
		log.Println("new development error", err)
		return err
	}
	zap.ReplaceGlobals(l)
	return nil
}

// InitProduction 将zap全局logger设置为生产
func InitProduction() error {
	l, err := zap.NewProduction()
	if err != nil {
		log.Println("new production error", err)
		return err
	}
	zap.ReplaceGlobals(l)
	return nil
}

// Sync 刷写全局logger缓存的日志，程序结束前调用
func Sync() {
	err := zap.L().Sync()
	if err != nil {
		log.Println("sync error", err)
	}
}
