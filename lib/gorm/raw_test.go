package gorm

import (
	"fmt"
	"testing"
)

func Test_raw_Exec(t *testing.T) {
	r, err := Crud.Raw.Exec("update users set age=19")
	fmt.Println(r, err)
}
