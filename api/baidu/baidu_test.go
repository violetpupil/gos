package baidu

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func Test_baidu_GetAccessToken(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
	err = InitEnv()
	if err != nil {
		panic(err)
	}

	r, err := Baidu.GetAccessToken()
	fmt.Println(r, err)
}
