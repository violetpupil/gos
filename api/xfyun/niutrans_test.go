package xfyun

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/lib/godotenv"
)

func Test_niuTrans_Translate(t *testing.T) {
	// 日志不加引号
	logrus.SetFormatter(&logrus.TextFormatter{DisableQuote: true})
	godotenv.Load("../../.env")
	InitEnv()
	text, err := Xfyun.NiuTrans.Translate("今天天气怎么样", LanguageCn, LanguageEn)
	fmt.Println(text, err)
}
