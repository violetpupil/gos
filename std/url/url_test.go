package url

import (
	"fmt"
	"testing"
)

func TestPath(t *testing.T) {
	fmt.Println(Path("https://pkg.go.dev/net/url"))
	fmt.Println(Path("/net/url"))
	fmt.Println(Path("net/url"))
}
