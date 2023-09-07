package baidu

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestGenAccessToken(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
	err = Init()
	if err != nil {
		panic(err)
	}

	err = GenAccessToken()
	fmt.Println(err)
}
