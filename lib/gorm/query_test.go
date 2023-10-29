package gorm

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/violetpupil/gos/lib/godotenv"
	"gorm.io/datatypes"
)

func Test_query_First(t *testing.T) {
	godotenv.Load("../../.env")
	err := InitMySQLEnv()
	if err != nil {
		panic(err)
	}
	err = Crud.AutoMigrate("", &User{})
	if err != nil {
		panic(err)
	}

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
	godotenv.Load("../../.env")
	err := InitMySQLEnv()
	if err != nil {
		panic(err)
	}
	err = Crud.AutoMigrate("", &User{})
	if err != nil {
		panic(err)
	}

	var tmp User
	r, err := Crud.R.FirstOrCreate(
		&tmp,
		User{Age: 10},
		User{Name: "jay"},
		User{Extend: datatypes.JSON("18888888888")},
	)
	fmt.Println(r, err)
}

func TestOrder(t *testing.T) {
	godotenv.Load("../../.env")
	err := InitMySQLEnv()
	if err != nil {
		panic(err)
	}
	err = Crud.AutoMigrate("", &User{})
	if err != nil {
		panic(err)
	}

	var users []User
	err = Order(D, "id").Limit(1).Find(&users).Error
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", users)
}
