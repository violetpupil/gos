package httptest

import "net/http/httptest"

type (
	// 测试用http.ResponseWriter
	ResponseRecorder = httptest.ResponseRecorder
)

var (
	// 创建ResponseRecorder
	NewRecorder = httptest.NewRecorder
)
