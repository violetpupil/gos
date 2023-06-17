package viper

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Unmarshal 加载配置文件到结构体
// name是文件名，可以不包含后缀
// t是文件类型 yaml
// path是文件所在目录
// s是结构体指针，标签用mapstructure
func Unmarshal(name, t, path string, s any) error {
	viper.SetConfigName(name)
	viper.SetConfigType(t)
	viper.AddConfigPath(path)
	// 加载配置文件
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorln("read in config error", err)
		return err
	}

	err = viper.Unmarshal(s)
	return err
}
