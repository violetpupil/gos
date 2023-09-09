package http

import "net/http"

// 响应体code
const (
	CodeSuccess       = 0 // 成功
	CodeArgumentError = 1 // 请求参数错误
	CodeUnknownError  = 2 // 未知错误
)

type (
	// ResponseWriter http响应构造器
	// WriteHeader() 写响应行和响应头
	// Write() 写响应体
	// 直接调用Write()，会自动调用WriteHeader()，使用200响应码
	// 如果响应头没有Content-Type，Write()会自动检测并设置
	ResponseWriter = http.ResponseWriter
	// Response http响应
	// Body 响应体
	// 读取 io.ReadAll(res.Body)
	// 客户端必须关闭 res.Body.Close()
	Response = http.Response
)
