// 当json解码到any接口时，接口实际类型对应关系
// []interface{}, for JSON arrays
// map[string]interface{}, for JSON objects
// https://pkg.go.dev/encoding/json#Unmarshal
//
// mapstructure可以直接将any再解码
// When decoding to a struct, mapstructure will use the field name by default to perform the mapping.
// 大小写不敏感
// The default struct tag that mapstructure looks for is "mapstructure"
// https://pkg.go.dev/github.com/mitchellh/mapstructure
package mapstructure
