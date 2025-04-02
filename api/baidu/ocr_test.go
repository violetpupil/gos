package baidu

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"testing"
)

func Test_baidu_OCR(t *testing.T) {
	_, err := Init("", "")
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
