package os

import (
	"fmt"
	"os"
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

func TestExist(t *testing.T) {
	r, err := Exist(".")
	fmt.Println(r, err)
}

func TestMkdir(t *testing.T) {
	err := Mkdir("tmp")
	fmt.Println(err)
}

func TestMkdirAll(t *testing.T) {
	err := MkdirAll("tmp/1/2")
	fmt.Println(err)
}

func TestRemove(t *testing.T) {
	err := os.Remove("tmp.txt")
	fmt.Println(err)
}

func TestRemoveAll(t *testing.T) {
	err := os.RemoveAll("tmp.txt")
	fmt.Println(err)
}
