# lib

## [decimal](https://pkg.go.dev/github.com/shopspring/decimal)

Arbitrary-precision fixed-point decimal numbers in go.

```golang
// 四舍五入
v1, exact := decimal.NewFromFloat(9.824).Round(2).Float64()
```

## [filenamify](https://pkg.go.dev/github.com/go-dora/filenamify)

Convert a string to a valid safe filename

```golang
filenamify.FilenamifyMustCompile
```
