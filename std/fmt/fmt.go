// 实现 error 或 fmt.Stringer 接口，会使 %+v 失效
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
