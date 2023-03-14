// 读写srt字幕文件
// https://docs.fileformat.com/zh/video/srt/
// github.com/wargarblgarbl/libgosubs/srt
package srt

import (
	"fmt"
	"os"
	"strings"

	"github.com/violetpupil/components/lib/logrus"
)

// Subtitle 单条字幕
type Subtitle struct {
	Id    int      // 字幕id，从1开始
	Start string   // 开始时间 00:00:00,000
	End   string   // 结束时间 00:00:00,000
	Line  []string // 多行字幕
}

// String 字幕字符串
func (sub Subtitle) String() string {
	lines := strings.Join(sub.Line, "\n")
	sen := fmt.Sprintf("%d\n%s --> %s\n%s", sub.Id, sub.Start, sub.End, lines)
	return sen
}

// Srt 返回srt字幕
func Srt(subs []Subtitle) string {
	var sens []string
	for _, sub := range subs {
		sens = append(sens, sub.String())
	}
	res := strings.Join(sens, "\n\n")
	return res
}

// WriteSrt 写srt字幕文件
// 文件不存在，创建文件，权限设置为可读可写
// 文件存在，清空文件，权限不变
func WriteSrt(subs []Subtitle, name string) error {
	f, err := os.Create(name)
	if err != nil {
		logrus.Error("create error ", err)
		return err
	}
	s := Srt(subs)
	_, err = fmt.Fprint(f, s)
	return err
}
