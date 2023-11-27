// https://req.cool/
package req

import "github.com/imroc/req/v3"

type (
	// 客户端
	// R() 创建请求
	// SetProxyURL() 设置代理
	Client = req.Client
	// 请求
	// Get() 发起get请求
	// SetHeaders() 设置请求头
	Request = req.Request
	// 响应
	// GetStatusCode() 获取响应码
	Response = req.Response
)

var (
	// 创建客户端
	C = req.C
)
