package url

import (
	"fmt"
	"net/url"
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

func TestParse(t *testing.T) {
	u, err := url.Parse("http://host?a=b")
	fmt.Println(u, err)
}
