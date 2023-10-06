// excel结构体绑定
package exl

import (
	"github.com/go-the-way/exl"
)

var (
	ReadFile = exl.ReadFile[exl.ReadConfigurator]
	// 写到指定文件
	// int类型字段 excel可能显示为科学计数法 用string类型可以避免
	Write = exl.Write[exl.WriteConfigurator]
	// 写到io.Writer
	WriteTo = exl.WriteTo[exl.WriteConfigurator]
)
