package reflect

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	var n int
	p := &n
	fmt.Println(Pointer(p))
}

type Tmp struct {
	A int
}

func TestFieldNames(t *testing.T) {
	r, err := FieldNames(Tmp{})
	fmt.Println(r, err)
}
