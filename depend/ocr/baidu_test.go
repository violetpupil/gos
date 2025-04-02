package ocr

import (
	"encoding/base64"
	"os"
	"testing"

	"github.com/violetpupil/gos/misc"
)

func TestBaiduOCR(t *testing.T) {
	bs, err := os.ReadFile("./test_data/test.jpg")
	if err != nil {
		t.Fatal(err)
	}
	str := base64.StdEncoding.EncodeToString(bs)

	r, err := BaiduOCR("", "", "", str)
	if err != nil {
		t.Error(err)
	} else {
		misc.JSONPrint(r)
	}
}
