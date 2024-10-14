package oss

import (
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/violetpupil/gos/misc"
)

var b *oss.Bucket

func TestMain(m *testing.M) {
	misc.InitDevelopmentLog()

	var err error
	b, err = Bucket("", "", "", "")
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestTmp(t *testing.T) {}
