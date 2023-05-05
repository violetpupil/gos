package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Res json响应
type Res struct {
	Code int    `json:"code"`           // 处理结果码 0成功 其它失败
	Msg  string `json:"msg"`            // 处理结果描述
	Data any    `json:"data,omitempty"` // 处理成功后返回的数据
}

// ResOK 响应成功
func ResOK(c *gin.Context, data any) {
	var res Res
	res.Code = 0
	res.Msg = "success"
	res.Data = data
	c.JSON(http.StatusOK, res)
}

// ResBadRequest 请求参数无效
func ResBadRequest(c *gin.Context) {
	var res Res
	res.Code = 1
	res.Msg = "parameter check failure"
	c.JSON(http.StatusOK, res)
}
