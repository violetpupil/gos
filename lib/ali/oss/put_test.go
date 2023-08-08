package oss

import (
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/violetpupil/gos/lib/godotenv"
)

func TestPutObjectFromFile(t *testing.T) {
	godotenv.Load("../../.env")
	err := InitEnv()
	if err != nil {
		panic(err)
	}
	PutObjectFromFile("test/oss.go", "./oss.go")
}

func TestPutObjectFromFileACL(t *testing.T) {
	godotenv.Load("../../.env")
	err := InitEnv()
	if err != nil {
		panic(err)
	}
	PutObjectFromFileACL("test/oss.go", "./oss.go", oss.ACLPublicRead)
}
