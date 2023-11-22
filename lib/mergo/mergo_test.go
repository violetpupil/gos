package mergo

import (
	"fmt"
	"testing"

	"github.com/imdario/mergo"
)

type S struct {
	A string
	B string
}

func TestMerge(t *testing.T) {
	a := S{A: "a"}
	b := S{B: "b"}
	err := mergo.Merge(&b, a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", b)
}
