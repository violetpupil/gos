package strconv

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/violetpupil/gos/spec/types"
)

var ParseInt = strconv.ParseInt

// ToString 将其它类型转为string，不支持的类型返回空字符串
func ToString(i any) string {
	switch v := i.(type) {
	case bool:
		return strconv.FormatBool(v)
	// 10进制表示
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	// 不用指数表示，返回全部小数位
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	default:
		return ""
	}
}

// UnitInt 将数字用千位表示 K M B
// 19000 -> 19K
// 19500 -> 19.5K
func UnitInt(n int) string {
	var s string
	if n < int(math.Pow10(3)) {
		s = strconv.Itoa(n)
	} else if n < int(math.Pow10(6)) {
		s = fmt.Sprintf("%d.%dK", n/int(math.Pow10(3)), types.Hundred(n))
	} else if n < int(math.Pow10(9)) {
		s = fmt.Sprintf("%d.%dM", n/int(math.Pow10(6)), types.HundredThousand(n))
	} else {
		s = fmt.Sprintf("%d.%dB", n/int(math.Pow10(9)), types.HundredMillion(n))
	}
	return strings.ReplaceAll(s, ".0", "")
}
