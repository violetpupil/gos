package otto

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	v, err := Run("1+1")
	fmt.Println(v, err)
}
