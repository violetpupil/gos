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

`OutputPaths` is a list of URLs or file paths to write logging output to. `c.OutputPaths = append(c.OutputPaths, "tmp.log")`

`ErrorOutputPaths` is a list of URLs to write internal logger errors to.
