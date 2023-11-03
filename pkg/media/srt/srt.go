// 读写srt字幕文件
// https://docs.fileformat.com/zh/video/srt/
// https://github.com/wargarblgarbl/libgosubs/tree/master/srt
package srt

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	// 单条字幕文本行数小于3
	ErrSingleLinesShort = errors.New("single lines too short")
	// 时间格式不对
	ErrTimeFormat = errors.New("time format incorrect")
)

// Subtitle 单条字幕
type Subtitle struct {
	Id    int      `json:"ed"`    // 字幕id，从1开始
	Start string   `json:"start"` // 开始时间 00:00:00,000
	End   string   `json:"end"`   // 结束时间 00:00:00,000
	Lines []string `json:"lines"` // 多行字幕
}

// NewSubtitle 将单条字幕文本转为结构体
func NewSubtitle(lines []string) (*Subtitle, error) {
	if len(lines) < 3 {
		return nil, ErrSingleLinesShort
	}

	s := new(Subtitle)
	id, err := strconv.Atoi(lines[0])
	if err != nil {
		logrus.Errorln("parse id error", err)
		return nil, err
	}
	s.Id = id

	times := strings.Split(lines[1], " --> ")
	if len(times) < 2 {
		return nil, ErrTimeFormat
	}
	s.Start = times[0]
	s.End = times[1]
	s.Lines = lines[2:]
	return s, nil
}

// String 单条字幕文本
func (sub *Subtitle) String() string {
	lines := strings.Join(sub.Lines, "\n")
	sen := fmt.Sprintf("%d\n%s --> %s\n%s", sub.Id, sub.Start, sub.End, lines)
	return sen
}

// SrtString 将字幕数组转为字幕文本
func SrtString(subs []*Subtitle) string {
	var sens []string
	for _, sub := range subs {
		sens = append(sens, sub.String())
	}
	res := strings.Join(sens, "\n\n")
	return res
}

// SrtSlice 从r读取字幕文本，转为字幕数组
func SrtSlice(r io.Reader) ([]*Subtitle, error) {
	subs := make([]*Subtitle, 0)
	// 单条字幕文本
	var seg []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		// srt字幕以空行分段
		if line != "" {
			seg = append(seg, line)
		} else {
			sub, err := NewSubtitle(seg)
			if err != nil {
				logrus.Errorln("new subtitle error", err)
				return nil, err
			}
			subs = append(subs, sub)
			seg = nil
		}
	}

	// 因为异常扫描终止
	if err := sc.Err(); err != nil {
		logrus.Errorln("scan stop error", err)
		return nil, err
	}
	return subs, nil
}

// WriteSrt 写srt字幕文件
// 文件不存在，创建文件，权限设置为可读可写
// 文件存在，清空文件，权限不变
func WriteSrt(subs []*Subtitle, name string) error {
	f, err := os.Create(name)
	if err != nil {
		logrus.Errorln("create error", err)
		return err
	}
	s := SrtString(subs)
	_, err = fmt.Fprint(f, s)
	return err
}

// LoadSrt 读取srt字幕文件
func LoadSrt(name string) ([]*Subtitle, error) {
	bs, err := os.ReadFile(name)
	if err != nil {
		logrus.Errorln("read file error", err)
		return nil, err
	}

	subs, err := ParseSrt(string(bs))
	return subs, err
}

// ParseSrt 解析srt字幕文本
func ParseSrt(s string) ([]*Subtitle, error) {
	reader := strings.NewReader(s)
	subs, err := SrtSlice(reader)
	return subs, err
}

// ParseSrtBytes 解析srt字幕字节串
func ParseSrtBytes(s []byte) ([]*Subtitle, error) {
	reader := bytes.NewReader(s)
	subs, err := SrtSlice(reader)
	return subs, err
}
