package xfyun

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/lib/godotenv"
	"github.com/violetpupil/components/std/json"
)

func Test_xfyun_Upload(t *testing.T) {
	godotenv.Load("../../.env")
	InitEnv()
	// 官方测试数据
	_, err := Xfyun.Upload("D:/file/跨境直播调研/测试文件/lfAsr_涉政.wav")
	fmt.Println(err)
}

func Test_xfyun_GetResult(t *testing.T) {
	// 日志不加引号
	logrus.SetFormatter(&logrus.TextFormatter{DisableQuote: true})
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

func Test_xfyun_OrderResult(t *testing.T) {
	logrus.SetReportCaller(true)

	var res GetResultRes
	err := json.Load("./test_data/get_result_res.json", &res)
	if err != nil {
		panic(err)
	}
	err = Xfyun.OrderResult(&res)
	fmt.Println(err)
}

func Test_xfyun_VideoTime(t *testing.T) {
	fmt.Println(Xfyun.VideoTime("2650"))
}
