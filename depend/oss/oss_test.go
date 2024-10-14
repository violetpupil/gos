package oss

import (
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var b *oss.Bucket

func TestMain(m *testing.M) {
	var err error
	b, err = Bucket("", "", "", "")
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestTmp(t *testing.T) {}
