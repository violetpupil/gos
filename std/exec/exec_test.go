package exec

import (
	"fmt"
	"testing"
)

func TestRunCommand(t *testing.T) {
	r, err := RunCommand("ipconfig")
	fmt.Println(r, err)
}
