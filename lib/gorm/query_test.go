package gorm

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/violetpupil/components/lib/godotenv"
)

func Test_query_First(t *testing.T) {
	godotenv.Load("../../.env")
	err := InitMySQLEnv()
	if err != nil {
		panic(err)
	}
	err = Crud.AutoMigrate("", &Tmp{})
	if err != nil {
		panic(err)
	}

	var tmp Tmp
	_, err = Crud.R.First(&tmp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", tmp)

	var extend Extend
	err = json.Unmarshal(tmp.Extend, &extend)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", extend)
}