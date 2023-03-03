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
