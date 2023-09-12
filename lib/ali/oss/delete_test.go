package oss

import (
	"fmt"
	"testing"

	"github.com/violetpupil/gos/lib/godotenv"
)

func TestDeleteObjects(t *testing.T) {
	godotenv.Load("../../../.env")
	err := InitEnv()
	if err != nil {
		panic(err)
	}

	err = DeleteObjects("tmp")
	fmt.Println(err)
}
