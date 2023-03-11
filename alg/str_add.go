package alg

import (
	"errors"
	"sort"

	"github.com/sirupsen/logrus"
)

// StrAdd 两个数字字符串相加
// 0-9 对应编码 48-57
func StrAdd(str1, str2 string) (string, error) {
	if str1 == "" || str2 == "" {
		return "", ErrNan
	}
	// 第一个字符串位数最多
	if len(str2) > len(str1) {
		str1, str2 = str2, str1
	}

	// 是否进位
	var carry bool
	var result []byte
	for i := 1; i <= len(str1); i++ {
		// 取对应位置的值，如果第二个字符串长度不够，取值0
		b1 := str1[len(str1)-i]
		var b2 byte = '0'
		if i <= len(str2) {
			b2 = str2[len(str2)-i]
		}

		// 计算当前位置的值，以及是否要进位
		var pos byte
		var err error
		pos, carry, err = PosAdd(b1, b2, carry)
		if err != nil {
			logrus.Error("pos add error ", err)
			return "", err
		}
		result = append(result, pos)
	}

	// 反转位置
	Reverse(result)
	return string(result), nil
}

// Reverse 反转切片
func Reverse(x any) {
	sort.Slice(x, func(i, j int) bool { return true })
}

// ErrNan 传入的参数不是数字
var ErrNan = errors.New("not a number")

// BtoN 编码值减去最小编码值，转为数字
func BtoN(b byte) (byte, error) {
	n := b - '0'
	if n > 9 {
		return 0, ErrNan
	}
	return n, nil
}

// PosAdd 两个数字字符串对应位置相加，carry表示上一次相加是否进位
// 返回当前位置相加结果，以及下一次是否进位
func PosAdd(b1, b2 byte, carry bool) (byte, bool, error) {
	// 转为数字
	n1, err := BtoN(b1)
	if err != nil {
		logrus.Error("n1 to number error ", err)
		return 0, false, err
	}
	n2, err := BtoN(b2)
	if err != nil {
		logrus.Error("n2 to number error ", err)
		return 0, false, err
	}

	sum := n1 + n2
	if carry {
		sum += 1
	}
	div, rem := sum/10, sum%10
	// 还原成编码值
	if div == 1 {
		return rem + '0', true, nil
	} else {
		return rem + '0', false, nil
	}
}
