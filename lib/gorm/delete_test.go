package gorm

import (
	"fmt"
	"testing"

	"github.com/violetpupil/gos/lib/godotenv"
)

func Test_delete_Delete(t *testing.T) {
	godotenv.Load("../../.env")
	err := InitMySQLEnv()
	if err != nil {
		panic(err)
	}
	err = Crud.AutoMigrate("", &User{})
	if err != nil {
		panic(err)
	}

	r, err := Crud.D.Delete(&User{Id: 1})
	fmt.Println(r, err)
}
