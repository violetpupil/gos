package regexp

import (
	"fmt"
	"testing"
)

func TestFindAllString(t *testing.T) {
	r, e := FindAllString("a", "abc")
	fmt.Println(r, e)
}

func TestFindAllStringSubmatch(t *testing.T) {
	r, e := FindAllStringSubmatch("a(b)", "abc")
	fmt.Println(r, e)
}
