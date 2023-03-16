package godotenv

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Load 从文件中加载环境变量
// 默认加载当前路径.env文件
// 不会覆盖已有的环境变量
func Load(filenames ...string) {
	err := godotenv.Load(filenames...)
	if err != nil {
		logrus.Error("godotenv load error ", err)
	}
}

// LoadParse 从文件中加载环境变量到结构体指针
// 默认加载当前路径.env文件
// 不会覆盖已有的环境变量
func LoadParse(v any, filenames ...string) error {
	Load(filenames...)
	err := env.Parse(v)
	return err
}
