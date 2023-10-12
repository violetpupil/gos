// 实现 error 或 fmt.Stringer 接口
// 打印时优先调用error.Error()
// 然后调用fmt.Stringer.String()
// 会使格式动词失效
//
// 格式动词
// Integer
// %c the character represented by the corresponding Unicode code point
package fmt

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

var (
	// 写io.Writer
	Fprintf = fmt.Fprintf
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
