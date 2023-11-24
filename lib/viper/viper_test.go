package viper

import (
	"fmt"
	"testing"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type C struct {
	Mode string
	Flag int
}

func TestWatchConfig(t *testing.T) {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var c C
	err = viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", c)

	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("config change %+v\n", in)
		err := viper.Unmarshal(&c)
		if err != nil {
			fmt.Println("config change fail", err)
		} else {
			fmt.Printf("config change success %+v\n", c)
		}
	})
	viper.WatchConfig()
	select {}
}

func TestUnmarshal(t *testing.T) {
	var c C
	err := Unmarshal("config.yaml", &c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", c)
}
