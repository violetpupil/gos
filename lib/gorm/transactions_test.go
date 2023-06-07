package gorm

import (
	"fmt"
	"testing"

	"github.com/violetpupil/components/lib/godotenv"
	"gorm.io/gorm"
)

func Test_transaction_Transaction(t *testing.T) {
	godotenv.Load("../../.env")
	err := InitMySQLEnv()
	if err != nil {
		panic(err)
	}

	err = Crud.T.Transaction(func(tx *gorm.DB) error {
		return nil
	})
	fmt.Println(err)
}
