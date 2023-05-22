package wsg

import (
	"fmt"
	"testing"
)

func TestChat(t *testing.T) {
	err := Chat("ws://localhost:8080/ws")
	fmt.Println(err)
}
