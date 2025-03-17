# [zap](https://pkg.go.dev/go.uber.org/zap)

```golang
// panic日志
defer func() {
 if err := recover(); err != nil {
  log.Error("panic occupy", zap.Any("err", err), zap.Stack("stack"))
 }
}()
```

## Config

`OutputPaths` is a list of URLs or file paths to write logging output to.

```golang
// 追加写入文件
file := fmt.Sprintf("tmp.%s.log", time.Now().Format("20060102150405"))
c.OutputPaths = append(c.OutputPaths, file)
```

`ErrorOutputPaths` is a list of URLs to write internal logger errors to.
