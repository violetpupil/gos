// 模板函数列表 https://masterminds.github.io/sprig/
package sprig

import sprig "github.com/Masterminds/sprig"

var (
	// html 模板函数映射
	FuncMap     = sprig.FuncMap
	HtmlFuncMap = sprig.HtmlFuncMap
	// text 模板函数映射
	TxtFuncMap = sprig.TxtFuncMap
)
