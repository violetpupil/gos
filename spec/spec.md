# [go](https://go.dev/)

[下载](https://go.dev/dl/)

[go文档](https://go.dev/doc/)

[Specification](https://go.dev/ref/spec)

[predeclared identifiers](https://pkg.go.dev/builtin)

## Operator

Unary operators have the highest precedence.

```text
Precedence    Operator
    5             *  /  %  <<  >>  &  &^
    4             +  -  |  ^
    3             ==  !=  <  <=  >  >=
    2             &&
    1             ||
```

## 函数

`copy()` 深拷贝切片，返回拷贝元素数=min(len(dst), len(src))
