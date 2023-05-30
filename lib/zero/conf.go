// 配置
// 直接从环境变量读取 `json:",env=SERVICE_NAME"`
// 或者配置yaml文件 ${SERVICE_NAME}
package zero

import "github.com/zeromicro/go-zero/core/conf"

var (
	// 加载配置，失败则返回错误
	Load = conf.Load
)

// MustLoad 加载配置，失败则退出程序
func MustLoad(path string, v any) {
	conf.MustLoad(path, v, conf.UseEnv())
}
