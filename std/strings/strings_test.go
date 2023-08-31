package strings

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	r := strings.Split("", ",")
	fmt.Println(len(r))
}
