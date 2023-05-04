package gin

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/violetpupil/components/lib/swag"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

// AdminGroup 管理页面
func AdminGroup(e *gin.Engine) {
	g := e.Group("/admin", AdminAuth)
	// 注册pprof处理器，提供运行时数据，路径以/admin/debug/pprof/开头
	pprof.RouteRegister(g)

	// 添加监控数据页面，访问路径为/metrics
	// 配置prometheus和grafana参考
	// https://grafana.com/oss/prometheus/exporters/go-exporter/
	p := ginprometheus.NewPrometheus("gin")
	p.Use(e)

	// swagger文档
	// 访问路径 /swagger/index.html
	// 必须导入生成的docs包
	swag.GinSwagger(e)
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
