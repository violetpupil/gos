package path

import (
	"fmt"
	"testing"
)

func TestFilename(t *testing.T) {
	fmt.Println(Filename("/a.txt"))
}
