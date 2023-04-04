package srt

import (
	"fmt"
	"testing"
)

func TestNewSubtitle(t *testing.T) {
	lines := []string{
		"1",
		"00:00:37,920 --> 00:00:40,120",
		"纽约市民们 暴风雨似乎已经过去",
		"It looks like the storm has passed.",
	}
	s, e := NewSubtitle(lines)
	fmt.Printf("%#v %v\n", s, e)
}

func TestLoadSrt(t *testing.T) {
	subs, err := LoadSrt("./test_data/john wick.srt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", subs[len(subs)-1])
}
