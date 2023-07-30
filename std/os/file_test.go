package os

import (
	"strings"
	"testing"
)

func TestRename(t *testing.T) {
	Rename("", func(s string) string {
		return strings.TrimPrefix(s, "")
	})
}
