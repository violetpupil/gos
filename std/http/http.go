package http

import "net/http"

const (
	MethodGet = http.MethodGet
)

const (
	HeaderContentType = "Content-Type"
)

// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
const (
	StatusOK                  = http.StatusOK
	StatusBadRequest          = http.StatusBadRequest       // 400
	StatusForbidden           = http.StatusForbidden        // 403
	StatusMethodNotAllowed    = http.StatusMethodNotAllowed // 405
	StatusInternalServerError = http.StatusInternalServerError
)

var (
	// 运行http服务，handler通常是nil
	// 默认地址 0.0.0.0:80
	ListenAndServe = http.ListenAndServe
)
