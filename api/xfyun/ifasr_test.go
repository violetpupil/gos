package xfyun

import (
	"fmt"
	"testing"

	"github.com/violetpupil/components/lib/godotenv"
)

func Test_xfyun_Upload(t *testing.T) {
	godotenv.Load("../../.env")
	InitEnv()
	_, err := Xfyun.Upload("D:/file/跨境直播调研/测试文件/lfAsr_涉政.wav")
	fmt.Println(err)
}

func Test_xfyun_GetResult(t *testing.T) {
	godotenv.Load("../../.env")
	InitEnv()
	_, err := Xfyun.GetResult("DKHJQ20230310113506626000968003F300000")
	fmt.Println(err)
}
