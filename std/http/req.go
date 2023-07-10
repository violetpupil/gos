package http

import "net/http"

// Request http请求
// Header 请求头
// Form 包含查询字符串和请求体中的表单数据，必须先调用ParseForm()
type Request = http.Request
