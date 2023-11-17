package gorm

import (
	"fmt"
	"testing"
)

func Test_delete_Delete(t *testing.T) {
	r, err := Crud.D.Delete(&User{Id: 1})
	fmt.Println(r, err)
}
