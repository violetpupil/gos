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

	os, err := ListObjects(path, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", os)
}
