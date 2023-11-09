package middle

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/violetpupil/gos/lib/gin/common"
)

// LogContext 记录请求响应信息中间件
// https://cloud.tencent.com/developer/article/2235645
func LogContext(c *gin.Context) {
	logger := logrus.WithFields(logrus.Fields{
		"ip":     c.ClientIP(),
		"method": c.Request.Method,
		"url":    c.Request.URL, // 包括查询字符串
	})

	// 读取请求体，并关闭，防止可能的泄露
	// 防止之后读取失败，填充新请求体
	reqBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Errorln("read all request body error", err)
		common.AbortBadRequest(c)
		return
	}
	c.Request.Body.Close()
	c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	logger = logger.WithField("reqBody", string(reqBody))
	logger.Infoln("request info")

	// 换掉响应构造器
	writer := &ResponseWriter{
		Body:           new(bytes.Buffer),
		ResponseWriter: c.Writer,
	}
	c.Writer = writer

	// 嵌套执行下一个处理器
	c.Next()

	status := c.Writer.Status()
	logger.WithFields(logrus.Fields{
		"status":  fmt.Sprintf("%d %s", status, http.StatusText(status)),
		"resBody": writer.Body.String(),
	}).Infoln("response info")
}
