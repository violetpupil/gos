package httptest

import (
	"net/http"
	"net/http/httptest"
)

type (
	// 测试用http.ResponseWriter
	ResponseRecorder = httptest.ResponseRecorder
)

// TestHandler 测试处理器函数
func TestHandler(handler http.HandlerFunc, req *http.Request) *http.Response {
	w := httptest.NewRecorder()
	handler(w, req)
	// 返回响应
	res := w.Result()
	return res
}
