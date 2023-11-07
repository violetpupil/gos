// 如果S嵌套T，T的方法可以通过S或*S调用，*T的方法只能通过*S调用
// 如果S嵌套*T，T和*T的方法都可以通过S或*S调用
package types

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
