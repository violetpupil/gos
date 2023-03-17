package http

import "net/http"

const (
	HeaderContentType = "Content-Type"
)

var (
	// 运行http服务，handler通常是nil
	// 默认地址 0.0.0.0:80
	ListenAndServe = http.ListenAndServe
)
