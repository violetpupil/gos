package resty

import (
	"fmt"
	"testing"
)

func TestSetProxy(t *testing.T) {
	SetProxy("http://proxyserver:8888")
	res, err := Get("https://lumtest.com/myip.json", nil)
	fmt.Println(res.String(), err)
}
