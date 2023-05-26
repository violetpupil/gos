package uuid

import (
	"fmt"
	"testing"
)

func TestNewRandom(t *testing.T) {
	r, err := NewRandom()
	fmt.Println(r, err)
}
