// excel结构体绑定
package exl

import (
	"github.com/go-the-way/exl"
)

var (
	ReadFile = exl.ReadFile[exl.ReadConfigurator]
	Write    = exl.Write[exl.WriteConfigurator]
)
