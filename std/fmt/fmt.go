// 实现 error 或 fmt.Stringer 接口，会使 %+v 失效
package fmt

import "fmt"

var (
	// 写io.Writer
	Fprintf = fmt.Fprintf
)
