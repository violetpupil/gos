package fastjson

import (
	"fmt"
	"testing"

	"github.com/valyala/fastjson"
)

func TestGetInt(t *testing.T) {
	bs := []byte(`{"id": 1}`)
	r := fastjson.GetInt(bs, "id")
	fmt.Println(r)
}
