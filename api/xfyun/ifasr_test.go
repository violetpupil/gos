package xfyun

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/violetpupil/components/lib/godotenv"
	"github.com/violetpupil/components/lib/logrus"
)

func Test_xfyun_Upload(t *testing.T) {
	godotenv.Load("../../.env")
	InitEnv()
	_, err := Xfyun.Upload("D:/file/跨境直播调研/测试文件/lfAsr_涉政.wav")
	fmt.Println(err)
}

func Test_xfyun_GetResult(t *testing.T) {
	// 日志不加引号
	logrus.Init()
	godotenv.Load("../../.env")
	InitEnv()
	_, err := Xfyun.GetResult("DKHJQ20230310113506626000968003F300000")
	fmt.Println(err)
}

func TestLattice_UnmarshalJSON(t *testing.T) {
	s := `{"json_1best": "{\"st\": {\"bg\": \"bg\"}}"}`
	var l Lattice
	err := json.Unmarshal([]byte(s), &l)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", l)
}
