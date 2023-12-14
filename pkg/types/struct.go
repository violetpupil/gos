// 如果S嵌套T，T的方法可以通过S或*S调用，*T的方法只能通过*S调用
// 如果S嵌套*T，T和*T的方法都可以通过S或*S调用
// 结构体标签 `key1:"value1" key2:"value2"`
package types

// KeyValue 键值对
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
