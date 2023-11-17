package gorm

import (
	"testing"

	"github.com/violetpupil/gos/lib/godotenv"
	"gorm.io/datatypes"
)

type Extend struct {
	Addr string
}

type User struct {
	Id   int
	Age  int    `gorm:"uniqueIndex:idx_age_name"`
	Name string `gorm:"uniqueIndex:idx_age_name;type:varchar(10)"`
	// 用字符串操作，用数字存储
	CompanyID string `gorm:"type:bigint"`
	Beauty    bool
	Extend    datatypes.JSON
}

func TestMain(m *testing.M) {
	godotenv.Load("../../.env")
	err := InitMySQLEnv()
	if err != nil {
		panic(err)
	}
	err = Crud.AutoMigrate("", &User{})
	if err != nil {
		panic(err)
	}

	m.Run()
}
