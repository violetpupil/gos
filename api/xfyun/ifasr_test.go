package xfyun

import (
	"fmt"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/lib/godotenv"
	"github.com/violetpupil/components/std/json"
)

func Test_xfyun_Upload(t *testing.T) {
	godotenv.Load("../../.env")
	InitEnv()
	// 官方测试数据
	_, err := Xfyun.LfAsr.Upload("D:/file/跨境直播调研/测试文件/lfAsr_涉政.wav")
	fmt.Println(err)
}

func Test_xfyun_GetResult(t *testing.T) {
	// 日志不加引号
	logrus.SetFormatter(&logrus.TextFormatter{DisableQuote: true})
	godotenv.Load("../../.env")
	InitEnv()
	_, err := Xfyun.LfAsr.GetResult("DKHJQ20230310113506626000968003F300000")
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

func Test_xfyun_OrderResult(t *testing.T) {
	logrus.SetReportCaller(true)
	InitEnv()

	var res GetResultRes
	err := json.Load("./test_data/get_result_res.json", &res)
	if err != nil {
		panic(err)
	}
	sub, err := Xfyun.LfAsr.OrderResult(&res)
	fmt.Println(sub, err)
}

func Test_xfyun_VideoTime(t *testing.T) {
	InitEnv()
	fmt.Println(Xfyun.LfAsr.VideoTime("2650"))
}

func Test_xfyun_SpeechToText(t *testing.T) {
	// 日志不加引号
	logrus.SetFormatter(&logrus.TextFormatter{DisableQuote: true})
	godotenv.Load("../../.env")
	InitEnv()
	// 官方测试数据
	sub, err := Xfyun.LfAsr.SpeechToText("D:/file/跨境直播调研/测试文件/lfAsr_涉政.wav", 3*time.Minute)
	fmt.Println(sub, err)
}
