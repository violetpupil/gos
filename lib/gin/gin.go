// https://github.com/gin-gonic/gin/blob/master/docs/doc.md
// https://gin-gonic.com/docs/examples/
// https://github.com/gin-gonic/examples
package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Run 运行http服务，直到错误发生，默认端口8080
func Run(addr ...string) {
	// Logger and Recovery middleware already attached
	e := gin.Default()
	Admin(e)

	// 业务组
	g := e.Group("", LogContext)
	g.GET("echo", func(c *gin.Context) {})
	fmt.Println(e.Run(addr...))
}
