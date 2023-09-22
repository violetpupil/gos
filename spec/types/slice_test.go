package types

import (
	"fmt"
	"testing"
)

func TestDiv(t *testing.T) {
	r := Div[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
	fmt.Println(r)
}
