package viper

import (
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// 设置配置文件改变时的处理函数
	OnConfigChange = viper.OnConfigChange
	// starts watching a config file for changes
	// 不会阻塞
	WatchConfig = viper.WatchConfig
)

// Unmarshal 加载配置文件到 s
// s 是结构体指针，标签用 mapstructure
func Unmarshal(in string, s any) error {
	viper.SetConfigFile(in)
	// 加载配置文件
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorln("read in config error", err)
		return err
	}

	err = viper.Unmarshal(s)
	return err
}

// UnmarshalWatch 加载配置文件到 s，并监听文件变动，同步到 s
// s 是结构体指针，标签用 mapstructure
func UnmarshalWatch(in string, s any) error {
	err := Unmarshal(in, s)
	if err != nil {
		logrus.Errorln("unmarshal error", err)
		return err
	}
	logrus.Infof("unmarshal config %+v", s)

	viper.OnConfigChange(func(in fsnotify.Event) {
		logger := logrus.WithField("event", in)
		err := viper.Unmarshal(s)
		if err != nil {
			logger.Errorln("config change fail", err)
		} else {
			logger.Infof("config change success %+v", s)
		}
	})
	viper.WatchConfig()
	return nil
}

// UnmarshalEnv 加载配置文件到结构体 自动加载环境变量
// f是配置文件路径
// t是文件类型 yaml
// s是结构体指针，标签用mapstructure
func UnmarshalEnv(f, t string, s any) error {
	bs, err := os.ReadFile(f)
	if err != nil {
		logrus.Errorln("read file error", err)
		return err
	}
	err = UnmarshalB(bs, t, s)
	return err
}

// UnmarshalB 加载配置字节串到结构体 自动加载环境变量
// t是格式类型 yaml
// s是结构体指针，标签用mapstructure
func UnmarshalB(b []byte, t string, s any) error {
	viper.SetConfigType(t)
	str := os.ExpandEnv(string(b))
	err := viper.ReadConfig(strings.NewReader(str))
	if err != nil {
		logrus.Errorln("read config error", err)
		return err
	}

	err = viper.Unmarshal(s)
	return err
}
