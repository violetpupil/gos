package oss

import (
	"fmt"
	"testing"

	"github.com/violetpupil/gos/lib/godotenv"
)

func TestManage_ListObjects(t *testing.T) {
	godotenv.Load("../../.env")
	err := InitEnv()
	if err != nil {
		panic(err)
	}
	os, err := ListObjects("tools")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", os)
}
