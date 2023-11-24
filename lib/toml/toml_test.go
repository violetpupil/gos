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
	O       struct{ A string }
	S       []struct{ B string }
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
	c.S = append(c.S, struct{ B string }{"s1"})
	c.S = append(c.S, struct{ B string }{"s2"})
	err := MarshalFile("tmp.toml", c)
	fmt.Println(err)
}
