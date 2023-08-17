package sha256

import (
	"fmt"
	"testing"
)

func TestSum256(t *testing.T) {
	r := Sum256("a")
	fmt.Println(r)
}
