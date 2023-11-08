package gorm

import (
	"fmt"
	"testing"

	"github.com/violetpupil/gos/lib/godotenv"
)

func Test_raw_Exec(t *testing.T) {
	godotenv.Load("../../.env")
	err := InitMySQLEnv()
	if err != nil {
		panic(err)
	}
	err = Crud.AutoMigrate("", &User{})
	if err != nil {
		panic(err)
	}

	r, err := Crud.Raw.Exec("update users set age=19")
	fmt.Println(r, err)
}
