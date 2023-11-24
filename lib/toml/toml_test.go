package toml

import (
	"fmt"
	"testing"
)

type MyConfig struct {
	Version int
	Name    string
	Tags    []string
}

func TestUnmarshalFile(t *testing.T) {
	var c MyConfig
	err := UnmarshalFile("data.toml", &c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", c)
}
