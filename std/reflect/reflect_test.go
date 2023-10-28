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

type Embed struct {
	B int
}

type Tmp struct {
	*Embed
	A int
}

func TestFieldNames(t *testing.T) {
	r, err := FieldNames(Tmp{Embed: &Embed{}})
	fmt.Println(r, err)
}
