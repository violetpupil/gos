package path

import (
	"fmt"
	"path"
	"testing"
)

func TestFilename(t *testing.T) {
	fmt.Println(Filename("/a.txt"))
}

func TestJoin(t *testing.T) {
	fmt.Println(path.Join("a", "b"))
}
