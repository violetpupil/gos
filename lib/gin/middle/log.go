package middle

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LogContext 记录请求响应信息中间件
func LogContext(c *gin.Context) {
	logger := logrus.WithFields(logrus.Fields{
		"method":    c.Request.Method,
		"path":      c.Request.URL, // 包括查询字符串
		"reqHeader": c.Request.Header,
		"reqBody":   nil,
	})
	logger.Infoln("request info")

	// 嵌套执行下一个处理器
	c.Next()

	status := c.Writer.Status()
	logger.WithFields(logrus.Fields{
		"status":    fmt.Sprintf("%d %s", status, http.StatusText(status)),
		"resHeader": nil,
		"resBody":   nil,
	}).Infoln("response info")
}
