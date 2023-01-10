package gorm

import (
	"errors"

	"gorm.io/gorm"
)

func RecordNotFound(e error) bool {
	return errors.Is(e, gorm.ErrRecordNotFound)
}
