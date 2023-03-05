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

func TestPost(t *testing.T) {
	Init(true)
	res, err := Post("https://httpbin.org/post", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestPostFile(t *testing.T) {
	Init(true)
	res, err := PostFile("https://httpbin.org/post", nil, "./method.go")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res)
}
