package oss

import "testing"

func TestInitToken(t *testing.T) {
	err := InitToken("")
	if err != nil {
		panic(err)
	}
}
