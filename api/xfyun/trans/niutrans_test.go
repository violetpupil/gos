package trans

import (
	"fmt"
	"testing"

	"github.com/caarlos0/env/v7"
	"github.com/sirupsen/logrus"
	"github.com/violetpupil/gos/lib/godotenv"
)

func Test_niuTrans_Translate(t *testing.T) {
	// 日志不加引号
	logrus.SetFormatter(&logrus.TextFormatter{DisableQuote: true})
	godotenv.Load("../../../.env")

	var a NiuTrans
	err := env.Parse(&a)
	if err != nil {
		panic(err)
	}
	text, err := a.Translate("今天天气怎么样", "cn", "en")
	fmt.Println(text, err)
}
