// 配置
// 直接从环境变量读取 `json:",env=SERVICE_NAME"`
// 或者yaml文件读取环境变量 ${SERVICE_NAME}
package zero

import "github.com/zeromicro/go-zero/core/conf"

var (
	// 加载配置，失败则退出程序
	MustLoad = conf.MustLoad
	// 加载配置，失败则返回错误
	Load = conf.Load
	// 使用环境变量
	UseEnv = conf.UseEnv
)
