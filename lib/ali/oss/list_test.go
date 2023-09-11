package oss

import (
	"fmt"
	"testing"

	"github.com/violetpupil/gos/lib/godotenv"
)

func TestListObjects(t *testing.T) {
	godotenv.Load("../../../.env")
	err := InitEnv()
	if err != nil {
		panic(err)
	}

	os, err := ListObjects("tools", true)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(os))
	for _, o := range os {
		fmt.Printf("%+v\n", o)
	}
}

func TestListObjectsMaxKeys(t *testing.T) {
	godotenv.Load("../../../.env")
	err := InitEnv()
	if err != nil {
		panic(err)
	}

	os, err := ListObjectsMaxKeys("tools", 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(os))
	fmt.Printf("%+v\n", os)
}

func TestListObjectsDir(t *testing.T) {
	godotenv.Load("../../../.env")
	err := InitEnv()
	if err != nil {
		panic(err)
	}

	dirs, err := ListObjectsDir("tools/")
	if err != nil {
		panic(err)
	}
	fmt.Println(dirs)
}
