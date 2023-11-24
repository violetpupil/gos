package viper

import (
	"fmt"
	"testing"

	"github.com/violetpupil/gos/lib/logrus"
)

type C struct {
	Mode string
	Flag int
}

func TestUnmarshal(t *testing.T) {
	var c C
	err := Unmarshal("config.yaml", &c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", c)
}

func TestUnmarshalWatch(t *testing.T) {
	logrus.DisableQuote()
	var c C
	err := UnmarshalWatch("config.yaml", &c)
	if err != nil {
		panic(err)
	}
	select {}
}
