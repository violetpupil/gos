package http

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandle(t *testing.T) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "", strings.NewReader("hi"))
	if err != nil {
		panic(err)
	}

	Handle(w, req)
	res := w.Result()
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Status)
	fmt.Println(string(body))
}
