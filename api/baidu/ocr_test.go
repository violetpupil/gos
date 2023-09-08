package baidu

import (
	"fmt"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func Test_baidu_OCR(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
	err = Init()
	if err != nil {
		panic(err)
	}

	r, err := Baidu.OCR("")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", r)
}
