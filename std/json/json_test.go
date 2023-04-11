package json

import "testing"

func TestStdout(t *testing.T) {
	Stdout(map[string]string{"a": "b"})
}
