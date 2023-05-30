package types

import (
	"fmt"
	"testing"
)

func TestToSlice(t *testing.T) {
	r := ToSlice([]string{"a", "b"})
	fmt.Println(r)
}
