package k8s

import (
	"fmt"
	"testing"

	"k8s.io/client-go/util/homedir"
)

func TestHomeDir(t *testing.T) {
	fmt.Println(homedir.HomeDir())
}
