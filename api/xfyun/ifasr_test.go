package xfyun

import (
	"fmt"
	"testing"

	"github.com/violetpupil/components/lib/godotenv"
)

func Test_xfyun_Upload(t *testing.T) {
	godotenv.Load()
	InitEnv()
	err := Xfyun.Upload("D:/file/跨境直播调研/测试文件/lfAsr_涉政.wav")
	fmt.Println(err)
}
