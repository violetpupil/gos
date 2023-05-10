package gorm

import (
	"fmt"
	"testing"

	"github.com/violetpupil/components/lib/godotenv"
)

func Test_delete_Delete(t *testing.T) {
	godotenv.Load("../../.env")
	err := InitMySQLEnv()
	if err != nil {
		panic(err)
	}
	err = Crud.AutoMigrate("", &Tmp{})
	if err != nil {
		panic(err)
	}

	r, err := Crud.D.Delete(&Tmp{Id: 1})
	fmt.Println(r, err)
}
