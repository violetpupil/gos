package oss

import (
	"fmt"
	"testing"

	"github.com/violetpupil/gos/lib/godotenv"
)

func TestGetObjectToFile(t *testing.T) {
	godotenv.Load("../../../.env")
	err := InitEnv()
	if err != nil {
		panic(err)
	}
	r, err := GetObjectToFile("tmp.txt", "tmp")
	fmt.Println(r, err)
}
