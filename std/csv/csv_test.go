package csv

import (
	"fmt"
	"testing"
)

func TestReadAll(t *testing.T) {
	r, e := ReadAll("../../api/ali/lang.csv")
	fmt.Println(r, e)
}
