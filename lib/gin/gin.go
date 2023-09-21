// https://github.com/gin-gonic/gin/blob/master/docs/doc.md
// https://gin-gonic.com/docs/examples/
// https://github.com/gin-gonic/examples
//
// 项目结构示例
// handler 处理器函数
// - req 请求对象
// - res 响应对象
package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/violetpupil/gos/lib/gin/middle"
)

type (
	// H is a shortcut for map[string]any
	H = gin.H
)

// Run 运行http服务，直到错误发生，默认端口8080
func Run(addr ...string) {
	// Logger and Recovery middleware already attached
	e := gin.Default()
	Admin(e)

	// 业务组
	g := e.Group("", middle.LogContext)
	g.Any("echo", func(c *gin.Context) {})
	fmt.Println(e.Run(addr...))
}
