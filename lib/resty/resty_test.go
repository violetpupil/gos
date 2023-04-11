package resty

import (
	"fmt"
	"testing"
)

func TestSetProxyPart(t *testing.T) {
	err := SetProxyPart(
		"http://proxyserver:8888",
		"user",
		"pass",
	)
	if err != nil {
		panic(err)
	}
	res, err := Get("https://lumtest.com/myip.json", nil)
	fmt.Println(res.String(), err)
}
