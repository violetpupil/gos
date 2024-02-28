# lib

## [decimal](https://pkg.go.dev/github.com/shopspring/decimal)

[Golang四舍五入保留两位小数](https://zhuanlan.zhihu.com/p/152050239)

```golang
v1, _ := decimal.NewFromFloat(9.824).Round(2).Float64()
v2, _ := decimal.NewFromFloat(9.826).Round(2).Float64()
v3, _ := decimal.NewFromFloat(9.8251).Round(2).Float64()
fmt.Println(v1, v2, v3)
```

## [filenamify](https://pkg.go.dev/github.com/go-dora/filenamify)

```golang
filenamify.FilenamifyMustCompile
```
