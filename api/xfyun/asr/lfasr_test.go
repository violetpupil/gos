package asr

import (
	"fmt"
	"testing"
	"time"

	"github.com/caarlos0/env/v7"
	"github.com/sirupsen/logrus"
	"github.com/violetpupil/gos/lib/godotenv"
	"github.com/violetpupil/gos/std/json"
)

func Test_xfyun_Upload(t *testing.T) {
	godotenv.Load("../../../.env")

	var a LfAsr
	err := env.Parse(&a)
	if err != nil {
		panic(err)
	}
	// 官方测试数据
	_, err = a.Upload("D:/file/跨境直播调研/测试文件/lfAsr_涉政.wav")
	fmt.Println(err)
}

func Test_xfyun_GetResult(t *testing.T) {
	// 日志不加引号
	logrus.SetFormatter(&logrus.TextFormatter{DisableQuote: true})
	godotenv.Load("../../../.env")

	var a LfAsr
	err := env.Parse(&a)
	if err != nil {
		panic(err)
	}
	_, err = a.GetResult("DKHJQ20230310113506626000968003F300000")
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

	var a LfAsr
	var res GetResultRes
	err := json.Load("./test_data/get_result_res.json", &res)
	if err != nil {
		panic(err)
	}
	sub, err := a.OrderResult(&res)
	fmt.Println(sub, err)
}

func Test_xfyun_VideoTime(t *testing.T) {
	var a LfAsr
	fmt.Println(a.VideoTime("2650"))
}

func Test_xfyun_SpeechToText(t *testing.T) {
	// 日志不加引号
	logrus.SetFormatter(&logrus.TextFormatter{DisableQuote: true})
	godotenv.Load("../../../.env")

	var a LfAsr
	err := env.Parse(&a)
	if err != nil {
		panic(err)
	}
	// 官方测试数据
	sub, err := a.SpeechToText("D:/file/跨境直播调研/测试文件/lfAsr_涉政.wav", 3*time.Minute)
	fmt.Println(sub, err)
}
