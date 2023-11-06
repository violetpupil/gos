package client

import (
	"fmt"
	"testing"
)

func TestGetUser(t *testing.T) {
	res, err := GetUser("")
	fmt.Println(res, err)
}
