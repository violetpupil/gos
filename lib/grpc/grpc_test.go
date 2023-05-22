package grpc

import (
	"fmt"
	"testing"
)

func TestServe(t *testing.T) {
	err := Serve("9090")
	fmt.Println(err)
}
