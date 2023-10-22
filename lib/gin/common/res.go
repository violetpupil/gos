package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 处理结果码
const (
	ResCodeOK                  = 0 // 处理成功
	ResCodeBadRequest          = 1 // 请求有误
	ResCodeInternalServerError = 2 // 服务有误
	ResCodeVerifyFail          = 3 // 验证失败 应明确错误原因
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
	res.Code = ResCodeOK
	res.Msg = "success"
	res.Data = data
	c.JSON(http.StatusOK, res)
}

// ResBadRequest 请求参数无效
func ResBadRequest(c *gin.Context) {
	var res Res
	res.Code = ResCodeBadRequest
	res.Msg = "parameter check failure"
	c.JSON(http.StatusOK, res)
}

// ResInternalServerError 服务错误
func ResInternalServerError(c *gin.Context) {
	var res Res
	res.Code = ResCodeInternalServerError
	res.Msg = "internal server error"
	c.JSON(http.StatusOK, res)
}

// AbortBadRequest 请求参数无效，放弃剩下的处理器
func AbortBadRequest(c *gin.Context) {
	var res Res
	res.Code = ResCodeBadRequest
	res.Msg = "parameter check failure"
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// ResVerifyFail 验证失败
func ResVerifyFail(c *gin.Context, msg string) {
	var res Res
	res.Code = ResCodeVerifyFail
	res.Msg = msg
	c.JSON(http.StatusOK, res)
}
