// 读写srt字幕文件
// https://docs.fileformat.com/zh/video/srt/
// https://github.com/wargarblgarbl/libgosubs/tree/master/srt
package srt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/violetpupil/components/lib/logrus"
)

// Subtitle 单条字幕
type Subtitle struct {
	Id    int      `json:"ed"`    // 字幕id，从1开始
	Start string   `json:"start"` // 开始时间 00:00:00,000
	End   string   `json:"end"`   // 结束时间 00:00:00,000
	Lines []string `json:"lines"` // 多行字幕
}

// String 字幕字符串
func (sub Subtitle) String() string {
	lines := strings.Join(sub.Lines, "\n")
	sen := fmt.Sprintf("%d\n%s --> %s\n%s", sub.Id, sub.Start, sub.End, lines)
	return sen
}

// SrtString 返回srt字幕
func SrtString(subs []Subtitle) string {
	var sens []string
	for _, sub := range subs {
		sens = append(sens, sub.String())
	}
	res := strings.Join(sens, "\n\n")
	return res
}

// SrtSlice 返回字幕数组
func SrtSlice(sub string) []Subtitle {
	return nil
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
	s := SrtString(subs)
	_, err = fmt.Fprint(f, s)
	return err
}

// ParseSrt 读取srt字幕文件
func ParseSrt(name string) ([]Subtitle, error) {
	f, err := os.Open(name)
	if err != nil {
		logrus.Errorln("open error", err)
		return nil, err
	}

	// 扫描每行
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	if scanner.Err() != nil {
		logrus.Errorln("scanner stop error", err)
	}
	return nil, nil
}
