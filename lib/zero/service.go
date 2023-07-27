package zero

import "github.com/zeromicro/go-zero/core/service"

// 运行环境
const (
	DevMode  = "dev"  // 开发
	TestMode = "test" // 测试
	RtMode   = "rt"   // 回归测试
	PreMode  = "pre"  // 预生产
	ProMode  = "pro"  // 生产
)

type (
	// ServiceConf is a service config.
	// Name service name
	// Mode 运行环境，默认为生产环境
	ServiceConf = service.ServiceConf
)
