package reflect

import (
	"fmt"
	"testing"
)

func TestElem(t *testing.T) {
	var n int
	p := &n
	fmt.Println(Elem(p))
}
