package gorm

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

func Test_transaction_Transaction(t *testing.T) {
	err := Crud.T.Transaction(func(tx *gorm.DB) error {
		return nil
	})
	fmt.Println(err)
}
