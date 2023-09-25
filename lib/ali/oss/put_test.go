package oss

import (
	"fmt"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/violetpupil/gos/lib/godotenv"
)

func TestPutObjectFromFile(t *testing.T) {
	godotenv.Load("../../../.env")
	err := InitEnv()
	if err != nil {
		panic(err)
	}
	err = PutObjectFromFile("test/oss.go", "./oss.go")
	fmt.Println(err)
}

func TestPutObjectFromFileACL(t *testing.T) {
	godotenv.Load("../../../.env")
	err := InitEnv()
	if err != nil {
		panic(err)
	}
	err = PutObjectFromFileACL("test/oss.go", "./oss.go", oss.ACLPublicRead)
	fmt.Println(err)
}

func TestPutObjectURL(t *testing.T) {
	godotenv.Load("../../../.env")
	err := InitEnv()
	if err != nil {
		panic(err)
	}
	err = PutObjectURL("tmp.jpg", "")
	fmt.Println(err)
}
