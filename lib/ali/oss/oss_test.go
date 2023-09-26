package oss

import (
	"fmt"
	"testing"
)

func TestInitToken(t *testing.T) {
	path, err := InitToken("")
	if err != nil {
		panic(err)
	}

	obj, err := First(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", obj)
}
