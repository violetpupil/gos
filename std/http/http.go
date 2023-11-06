package http

import "net/http"

const (
	HeaderContentType   = "Content-Type"
	HeaderContentLength = "Content-Length"
)

// 函数
var (
	// 运行http服务
	// handler 传 nil，使用 http.DefaultServeMux
	// 默认地址 0.0.0.0:80
	ListenAndServe = http.ListenAndServe
	// 创建路由器
	NewServeMux = http.NewServeMux
)

// 类型
type (
	// 处理器函数类型，并实现Handler接口
	HandlerFunc = http.HandlerFunc
	// 路由器
	ServeMux = http.ServeMux
)

// 默认路由器及操作函数
var (
	// 默认路由器
	DefaultServeMux = http.DefaultServeMux

	// 操作默认路由器
	// 注册处理器函数
	HandleFunc = http.HandleFunc
)
