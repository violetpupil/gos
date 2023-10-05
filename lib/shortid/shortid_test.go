package shortid

import (
	"fmt"
	"testing"

	"github.com/teris-io/shortid"
)

func TestGenerate(t *testing.T) {
	fmt.Println(shortid.Generate())
}
