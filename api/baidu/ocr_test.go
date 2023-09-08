package baidu

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func Test_baidu_OCR(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
	err = InitEnv()
	if err != nil {
		panic(err)
	}

	bs, err := os.ReadFile("./test_data/test.jpg")
	if err != nil {
		panic(err)
	}
	str := base64.StdEncoding.EncodeToString(bs)

	r, err := Baidu.OCR(str)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", r)
}
