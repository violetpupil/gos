# goroutine

## [ants](https://pkg.go.dev/github.com/panjf2000/ants/v2)

可以控制goroutine池的大小

```golang
var wg sync.WaitGroup

f := func(arg any) {
 defer wg.Done()
}

p, err := ants.NewPoolWithFunc(10, f)
if err != nil {
 log.Error("new pool error", zap.Error(err))
 return err
}
defer p.Release()

for _, id := range ids {
 wg.Add(1)
 p.Invoke(id)
}
 
wg.Wait()
```

## 并发执行

```golang
// 等待并发执行结束
var wg sync.WaitGroup
// 处理执行结果，避免加锁
receive := make(chan string, len(ids))
// 等待接收结束
waitReceive := make(chan int)
go func() {
 for id := range receive {
  fmt.Println(id)
 }
 waitReceive <- 0
}()
for _, id := range ids {
 wg.Add(1)
 go func(id string) {
  defer wg.Done()

  receive <- id
 }(id)
}
wg.Wait()
close(receive)
<-waitReceive
```
