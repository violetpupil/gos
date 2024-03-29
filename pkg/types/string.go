// 截取中文字符串
// string([]rune("你好")[:2])
package types

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Colors
const (
	Reset       = "\033[0m"    // 重置
	Red         = "\033[31m"   // 红
	Green       = "\033[32m"   // 绿
	Yellow      = "\033[33m"   // 黄
	Blue        = "\033[34m"   // 蓝
	Magenta     = "\033[35m"   // 紫
	Cyan        = "\033[36m"   // 淡蓝
	White       = "\033[37m"   // 白
	BlueBold    = "\033[34;1m" // 蓝加粗
	MagentaBold = "\033[35;1m" // 紫加粗
	RedBold     = "\033[31;1m" // 红加粗
	YellowBold  = "\033[33;1m" // 黄加粗
)

// Color 将文本上色，然后重置
func Color(c string, s string) string {
	return c + s + Reset
}

// Gen 生成连续数字字符串
func Gen(src string, max int) error {
	f, err := os.Create("tmp.txt")
	if err != nil {
		logrus.Errorln("create file error", err)
		return err
	}
	defer f.Close()

	for i := 1; i < max+1; i++ {
		fmt.Fprintf(f, src+"\n", i, i)
	}
	return nil
}
