package flag

import "flag"

var (
	// 传递布尔参数，调用函数
	BoolFunc = flag.BoolFunc
	// 返回指针
	Int = flag.Int
	// 传入指针
	IntVar = flag.IntVar
)
