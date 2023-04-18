package gorm

import (
	"fmt"
	"testing"

	"github.com/violetpupil/components/lib/godotenv"
)

func Test_create_Create(t *testing.T) {
	godotenv.Load("../../.env")
	err := InitMySQLEnv()
	if err != nil {
		panic(err)
	}
	err = Crud.AutoMigrate(&Tmp{})
	if err != nil {
		panic(err)
	}

	tmp := Tmp{Id: 1, Age: 1}
	c, e := Crud.C.Create(&tmp)
	fmt.Println(c, e)
}

func Test_create_CreateDoNothing(t *testing.T) {
	godotenv.Load("../../.env")
	err := InitMySQLEnv()
	if err != nil {
		panic(err)
	}
	err = Crud.AutoMigrate(&Tmp{})
	if err != nil {
		panic(err)
	}

	tmp := Tmp{Id: 1, Age: 1, Name: "1"}
	c, e := Crud.C.CreateDoNothing(&tmp)
	fmt.Println(c, e)
}