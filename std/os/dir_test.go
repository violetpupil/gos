package os

import (
	"fmt"
	"testing"
)

func TestMkdir(t *testing.T) {
	err := Mkdir("tmp")
	fmt.Println(err)
}

func TestMkdirAll(t *testing.T) {
	err := MkdirAll("tmp/1/2")
	fmt.Println(err)
}
