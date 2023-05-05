package gin

import (
	"net/url"

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
	// 解码查询字符串
	query, err := url.QueryUnescape(c.Request.URL.RawQuery)
	if err != nil {
		logrus.Errorln("query unescape error", err)
		AbortBadRequest(c)
		return
	}
	logrus.WithFields(logrus.Fields{
		"Query":  query,
		"Header": c.Request.Header,
	}).Infoln("request info")

	// 嵌套执行下一个处理器
	c.Next()
}
