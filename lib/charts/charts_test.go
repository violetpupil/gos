package charts

import (
	"fmt"
	"testing"
)

func TestNewBar(t *testing.T) {
	err := NewBar(
		"",
		[]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
		map[string][]any{
			"asc":  {1, 2, 3, 4, 5, 6, 7},
			"desc": {7, 6, 5, 4, 3, 2, 1},
		},
	)
	fmt.Println(err)
}
