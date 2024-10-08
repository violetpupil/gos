# [json](https://pkg.go.dev/encoding/json)

```golang
// 编码缩进两个空格
json.MarshalIndent(v, "", "  ")
// 解码
// 1. 字段名大小写不敏感
// 2. 当字段为any时
// json array -> []any
// json object > map[string]any
json.Unmarshal(data, v)
```

## 结构体标签

`json:"myName,omitempty"`

`omitempty` 当字段为默认值时，不编码

`string` 字段对应的json类型为string

## 第三方库

[json](https://pkg.go.dev/github.com/goccy/go-json) 替代标准库、快

[fastjson](https://pkg.go.dev/github.com/valyala/fastjson) 无需解码，直接获取字段值
