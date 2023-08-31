package os

import (
	"fmt"
	"strings"
	"testing"
)

func TestRename(t *testing.T) {
	Rename("", func(s string) string {
		return strings.TrimPrefix(s, "")
	})
}

func TestWriteFile(t *testing.T) {
	err := WriteFile("tmp.txt", nil)
	fmt.Println(err)
}
