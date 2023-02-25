package resty

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	res, err := Get("https://httpbin.org/get", false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", res)
}
