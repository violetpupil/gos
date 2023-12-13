package req

import (
	"fmt"
	"testing"

	"github.com/imroc/req/v3"
)

func TestResponse(t *testing.T) {
	res, err := req.C().R().Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	fmt.Println(res.GetStatus())
}
