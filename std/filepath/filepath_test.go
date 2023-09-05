package filepath

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestJoin(t *testing.T) {
	fmt.Println(filepath.Join("a", "b"))
}

func TestWalk(t *testing.T) {
	err := Walk("D:/test/avatar")
	fmt.Println(err)
}
