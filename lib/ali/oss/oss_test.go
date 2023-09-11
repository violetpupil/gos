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

	os, err := ListObjects(path, 1)
	if err != nil {
		panic(err)
	}
	for _, o := range os {
		fmt.Printf("%+v\n", o)
	}
}
