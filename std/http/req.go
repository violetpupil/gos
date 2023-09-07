package http

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Request http请求
// Header 请求头
// Body 请求体
// Form 包含查询字符串和请求体中的表单数据，必须先调用ParseForm()
type Request = http.Request

// Handle 服务处理函数
func Handle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorln("read request body error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logrus.Infoln("request body", string(body))
}
