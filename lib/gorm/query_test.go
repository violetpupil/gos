package gorm

import (
	"encoding/json"
	"fmt"
	"testing"

	"gorm.io/datatypes"
)

func Test_query_First(t *testing.T) {
	var tmp User
	ok, err := Crud.R.First(&tmp)
	if err != nil {
		panic(err)
	}
	if !ok {
		fmt.Println("not found")
		return
	}
	fmt.Printf("%+v\n", tmp)

	if len(tmp.Extend) == 0 {
		return
	}
	var extend Extend
	err = json.Unmarshal(tmp.Extend, &extend)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", extend)
}

func Test_query_FirstOrCreate(t *testing.T) {
	var tmp User
	r, err := Crud.R.FirstOrCreate(
		&tmp,
		User{Age: 10},
		User{Name: "jay"},
		User{Extend: datatypes.JSON("18888888888")},
	)
	fmt.Println(r, err)
}
