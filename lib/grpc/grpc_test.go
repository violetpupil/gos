package grpc

import (
	"fmt"
	"testing"
)

func TestServe(t *testing.T) {
	err := Serve("9090", nil)
	fmt.Println(err)
}

func TestDial(t *testing.T) {
	conn, err := Dial("localhost:9090")
	if err != nil {
		panic(err)
	}
	conn.Close()
}
