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

type (
	// 日志器
	// With() 添加字段
	Logger = zap.Logger
)

// InitDevelopment 将zap全局logger设置为开发
func InitDevelopment() error {
	c := zap.NewDevelopmentConfig()
	// 不输出调用栈
	c.DisableStacktrace = true
	l, err := c.Build()
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

// NewDevelopmentFile 使用预设开发配置创建logger，日志写到文件
// file指定文件，设置为空则使用zap.log
// stderr是否写到标准错误
// 可以使用os操作文件
func NewDevelopmentFile(file string, stderr bool, options ...zap.Option) (*zap.Logger, error) {
	if file == "" {
		file = "zap.log"
	}

	c := zap.NewDevelopmentConfig()
	// 不输出调用栈
	c.DisableStacktrace = true
	if stderr {
		c.OutputPaths = append(c.OutputPaths, file)
		c.ErrorOutputPaths = append(c.OutputPaths, file)
	} else {
		c.OutputPaths = []string{file}
		c.ErrorOutputPaths = []string{file}
	}
	return c.Build(options...)
}
