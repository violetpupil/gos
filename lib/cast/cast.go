package cast

import "github.com/spf13/cast"

var (
	// 尝试转为字符串，失败返回空字符串
	ToString = cast.ToString
	// 尝试转为字符串，失败返回空字符串和错误
	ToStringE = cast.ToStringE
	// 将其它类型切片转为int切片
	ToIntSlice = cast.ToIntSlice
)
