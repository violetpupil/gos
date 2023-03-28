package oss

import (
	"testing"

	"github.com/violetpupil/components/lib/godotenv"
)

func TestPutObjectFromFile(t *testing.T) {
	godotenv.Load("../../.env")
	InitEnv()
	PutObjectFromFile("test/oss.go", "./oss.go")
}
