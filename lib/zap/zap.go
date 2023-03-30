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
