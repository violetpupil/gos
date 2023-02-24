package sonyflake

import (
	"fmt"
	"testing"
)

func TestNextId(t *testing.T) {
	err := Init()
	if err != nil {
		panic(err)
	}
	fmt.Println(NextId())
}
