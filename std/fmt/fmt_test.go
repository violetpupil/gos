package fmt

import (
	"fmt"
	"testing"
)

type S struct{}

func (s S) String() string {
	return "s"
}

func TestPrintf(t *testing.T) {
	fmt.Printf("%s\n", S{})
}
