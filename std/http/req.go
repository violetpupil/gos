package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	MethodGet = http.MethodGet
)

// Request http请求
// Header 请求头
// Body 请求体
// Form 包含查询字符串和请求体中的表单数据，必须先调用ParseForm()
// FormValue() 获取表单值，请求体优先于查询字符串
type Request = http.Request

// Handle 服务处理函数
func Handle(w http.ResponseWriter, r *http.Request) {
	res := map[string]any{
		"code": CodeSuccess,
		"msg":  "success",
		"data": "",
	}

	// 读请求体
	var body string
	bs, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorln("read request body error", err)
		res["code"] = CodeArgumentError
		res["msg"] = "read request body error"
		goto Fin
	}
	body = string(bs)
	logrus.Infoln("read request body result", body)

Fin:
	bs, err = json.Marshal(res)
	if err != nil {
		logrus.Errorln("marshal response error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}
