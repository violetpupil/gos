package types

import (
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {
	err := Gen("%03d %03d", 999)
	fmt.Println(err)
}
