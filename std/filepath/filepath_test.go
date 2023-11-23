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
	err := Walk("video")
	fmt.Println(err)
}

func TestGlobDir(t *testing.T) {
	r, err := GlobDir("video")
	fmt.Println(r, err)
}
