package mergo

import "github.com/imdario/mergo"

var (
	// 合并同类型结构体或 map
	// src 非空字段将填充到 dst 空字段上
	Merge = mergo.Merge
)
