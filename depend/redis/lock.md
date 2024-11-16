# [lock](https://pkg.go.dev/github.com/go-redsync/redsync/v4)

```golang
lock := global.Redsync.NewMutex(
 key,
 redsync.WithExpiry(time.Minute),
 redsync.WithTries(1000), // 默认重试32次，每次间隔rand(50ms, 250ms)
)
err := lock.LockContext(c)
if err != nil {
 log.Error("lock error", zap.String("key", key), zap.Error(err))
 return err
} else {
 log.Info("lock success", zap.String("key", key))
}
defer func() {
 r, err := lock.UnlockContext(c)
 log.Info("unlock result", zap.String("key", key), zap.Bool("status", r), zap.Error(err))
}()
```
