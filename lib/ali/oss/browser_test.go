package oss

import (
	"fmt"
	"testing"
)

func TestParseAuthToken(t *testing.T) {
	r, err := ParseAuthToken("")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", r)
}
