// https://github.com/gin-gonic/gin/blob/master/docs/doc.md
// https://gin-gonic.com/docs/examples/
// https://github.com/gin-gonic/examples
//
// 项目结构示例
// handler -> service -> dao
//
// api 外部接口调用
// global 全局对象及初始化
// - init.go 初始化
// dao 数据库操作
// dto 数据传输对象
// - query 查询对象 有时可直接使用响应对象
// - req 请求对象 前端处理64位整型会越界，改用字符串
// - res 响应对象 前端处理64位整型会越界，改用字符串
// handler 处理器函数 接口逻辑
// service 服务逻辑 整合 api, dao, dto
// 可将服务划分为不同包
// - err.go 错误类型
package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/violetpupil/gos/lib/gin/handler"
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

	e.GET("excel", handler.Excel)
	// 业务组 - 记录请求响应
	gLog := e.Group("", middle.LogContext)
	gLog.Any("echo", func(c *gin.Context) {})
	fmt.Println(e.Run(addr...))
}
