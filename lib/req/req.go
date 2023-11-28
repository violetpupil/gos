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
	// SetQueryParam() 设置查询字符串
	// SetSuccessResult() 成功响应体
	Request = req.Request
	// 响应
	// GetStatus() 获取响应码及描述
	// GetStatusCode() 获取响应码
	// String() 响应体字符串
	Response = req.Response
)

var (
	// 创建客户端
	C = req.C
)
