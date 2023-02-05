package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// main 运行http服务在端口8080，直到错误发生
func main() {
	e := gin.Default()
	AdminGroup(e)
	fmt.Print(e.Run())
}

// AdminGroup 管理页面
func AdminGroup(e *gin.Engine) {
	g := e.Group("/admin", AdminAuth)
	// 注册pprof处理器，提供运行时数据，路径以/admin/debug/pprof/开头
	pprof.RouteRegister(g)
}

// AdminAuth 管理页面认证中间件
// 使用 qs 参数 ?hypercube=assemble 进行认证
func AdminAuth(c *gin.Context) {
	if c.Query("hypercube") != "assemble" {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.Next()
}
