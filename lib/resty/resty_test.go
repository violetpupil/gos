package resty

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	Init(true)
	res, err := Get("https://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res)
}
