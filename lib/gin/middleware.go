package gin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Cors 跨域中间件
func Cors() gin.HandlerFunc {
	f := cors.Default()
	return f
}

// LogContext 记录请求响应信息中间件
func LogContext(c *gin.Context) {
	logrus.Info("request info")
	// 嵌套执行下一个处理器
	c.Next()
	logrus.Info("response info")
}
