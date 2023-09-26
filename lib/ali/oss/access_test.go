package oss

import (
	"fmt"
	"testing"

	"github.com/violetpupil/gos/lib/godotenv"
)

func TestSignURL(t *testing.T) {
	godotenv.Load("../../../.env")
	err := InitEnv()
	if err != nil {
		panic(err)
	}
	r, err := SignURL("tmp.txt", 60)
	fmt.Println(r, err)
}
