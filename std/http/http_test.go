package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestListenAndServe(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
	})
	// 测试访问地址
	err := ListenAndServe("", nil)
	fmt.Println(err)
}
