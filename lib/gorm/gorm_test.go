package gorm

import (
	"fmt"
	"testing"

	"github.com/violetpupil/components/lib/godotenv"
)

func TestInitMySQLEnv(t *testing.T) {
	godotenv.Load("../../.env")
	err := InitMySQLEnv()
	fmt.Println(err)
}

type Tmp struct {
	Id     int
	Age    int    `gorm:"uniqueIndex:idx_age_name"`
	Name   string `gorm:"uniqueIndex:idx_age_name;type:varchar(10)"`
	Beauty bool
}

func Test_crud_AutoMigrate(t *testing.T) {
	godotenv.Load("../../.env")
	err := InitMySQLEnv()
	if err != nil {
		panic(err)
	}

	err = Crud.AutoMigrate("", &Tmp{})
	fmt.Println(err)
}
