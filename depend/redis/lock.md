# [lock](https://pkg.go.dev/github.com/go-redsync/redsync/v4)

过期设置为此次操作可能的最长时间。重试次数和间隔根据想等多久来设置。

```golang
lock := global.Redsync.NewMutex(
 key,
 redsync.WithExpiry(30*time.Second),
 redsync.WithTries(600), // 默认重试32次，每次间隔rand(50ms, 250ms)
)
err := lock.LockContext(c)
if errors.Is(err, redsync.ErrFailed) {
 log.Info("failed to acquire lock", zap.String("key", key), zap.Error(err))
 return err
} else if err != nil {
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
