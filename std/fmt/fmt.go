// 实现 error 或 fmt.Stringer 接口，会先转为字符串
// 打印时优先调用error.Error()
// 然后调用fmt.Stringer.String()
//
// 格式动词
// %v the value in a default format
// %+v 结构体打印字段名，其它类型按 %v 处理
// Integer
// %c the character represented by the corresponding Unicode code point
package fmt

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

var (
	// 根据格式字符串，创建 error
	// 使用 %w 格式动词，包裹 error
	// 返回的 error 会实现 Unwrap 方法
	Errorf = fmt.Errorf
	// 写io.Writer
	Fprintf = fmt.Fprintf
	// 参数间加空格，最后加换行
	Fprintln = fmt.Fprintln
)

// PrettyPrintln 美化换行打印
// v必须可以json编码
func PrettyPrintln(v any) {
	bs, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		logrus.Errorln("json marshal indent error", err)
		fmt.Println(v)
	} else {
		fmt.Println(string(bs))
	}
}
