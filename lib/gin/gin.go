package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Run 运行http服务，直到错误发生，默认端口8080
func Run(addr ...string) {
	// Logger and Recovery middleware already attached
	e := gin.Default()
	AdminGroup(e)

	e.POST("/echo", Echo)
	fmt.Println(e.Run(addr...))
}

// Echo 打印请求信息
func Echo(c *gin.Context) {
	fmt.Println("echo request header")
	fmt.Println(c.Request.Header)
}
