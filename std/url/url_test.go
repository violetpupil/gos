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

func TestQuery(t *testing.T) {
	r, err := Query("", "")
	fmt.Println(r, err)
}
