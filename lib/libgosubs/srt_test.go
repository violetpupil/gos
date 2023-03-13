package libGoSubs

import (
	"fmt"
	"testing"

	"github.com/wargarblgarbl/libgosubs/srt"
)

func TestWriteSrt(t *testing.T) {
	subs := []srt.Subtitle{
		{
			Id:    1,
			Start: "00:00:01,000",
			End:   "00:00:10,000",
			Line:  []string{"timer"},
		},
	}
	err := WriteSrt(subs, "tmp.srt")
	fmt.Println(err)
}
