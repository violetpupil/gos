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

	r, err := ListObjects(Prefix("tools"))
	if err != nil {
		panic(err)
	}
	fmt.Println(len(r.Objects))
	for _, o := range r.Objects {
		fmt.Printf("%+v\n", o)
	}
	fmt.Println(len(r.CommonPrefixes))
	for _, p := range r.CommonPrefixes {
		fmt.Println(p)
	}
}
