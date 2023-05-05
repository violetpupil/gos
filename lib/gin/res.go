package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Res json响应
type Res struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

// ResOK 响应成功
func ResOK(c *gin.Context, data any) {
	var res Res
	res.Code = 0
	res.Msg = "success"
	res.Data = data
	c.JSON(http.StatusOK, res)
}
