package toml

import (
	"fmt"
	"testing"
	"time"
)

type MyConfig struct {
	Version int
	Name    string
	Tags    []string
	T       time.Time
}

func TestUnmarshalFile(t *testing.T) {
	var c MyConfig
	err := UnmarshalFile("data.toml", &c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", c)
}

func TestMarshalFile(t *testing.T) {
	var c MyConfig
	c.T = time.Now()
	err := MarshalFile("tmp.toml", c)
	fmt.Println(err)
}
