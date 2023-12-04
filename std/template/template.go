// Templates are executed by applying them to a data structure.
//
// 使用 {{}} 表示 actions
// {{.Field}} 使用字段
//
// 可以使用全局函数和模板中的函数
// 全局函数
// print An alias for fmt.Sprint
// printf An alias for fmt.Sprintf
package template

import "text/template"

var (
	// 创建模板
	New = template.New
)

type (
	// 模板
	// Funcs() 添加函数
	// Parse() 设置模板文本
	// Execute() 填充模板并输出
	Template = template.Template
)
