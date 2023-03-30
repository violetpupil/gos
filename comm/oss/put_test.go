package oss

import (
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/violetpupil/components/lib/godotenv"
)

func TestPutObjectFromFile(t *testing.T) {
	godotenv.Load("../../.env")
	InitEnv()
	PutObjectFromFile("test/oss.go", "./oss.go")
}

func TestPutObjectFromFileACL(t *testing.T) {
	godotenv.Load("../../.env")
	InitEnv()
	PutObjectFromFileACL("test/oss.go", "./oss.go", oss.ACLPublicRead)
}
