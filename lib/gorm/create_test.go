package gorm

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_create_Create(t *testing.T) {
	extend := Extend{Addr: "China"}
	bs, err := json.Marshal(extend)
	if err != nil {
		panic(err)
	}
	tmp := User{Id: 1, Age: 1, Extend: bs}
	c, e := Crud.C.Create(&tmp)
	fmt.Println(c, e)
}

func Test_create_CreateDoNothing(t *testing.T) {
	tmp := User{Id: 1, Age: 1, Name: "1"}
	c, e := Crud.C.CreateDoNothing(&tmp)
	fmt.Println(c, e)
}

func Test_create_CreateUpdateAll(t *testing.T) {
	tmp := User{Id: 1, Age: 1, Name: "1"}
	e := Crud.C.CreateUpdateAll(&tmp)
	fmt.Println(e)
}
