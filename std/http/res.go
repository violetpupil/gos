package http

import "net/http"

// ResponseWriter http响应构造器
// WriteHeader() 写响应行和响应头
// Write() 写响应体
// 直接调用Write()，会自动调用WriteHeader()，使用200响应码
// 如果响应头没有Content-Type，Write()会自动检测并设置
type ResponseWriter = http.ResponseWriter
