# [pool](https://pkg.go.dev/github.com/panjf2000/ants/v2)

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
