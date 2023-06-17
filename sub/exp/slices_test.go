package exp

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func TestContains(t *testing.T) {
	r := slices.Contains([]string{""}, "")
	fmt.Println(r)
}
