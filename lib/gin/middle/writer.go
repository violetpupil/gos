package middle

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

// ResponseWriter 响应体双写
type ResponseWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

// Write 重写 http.ResponseWriter.Write() 方法
func (w *ResponseWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}
