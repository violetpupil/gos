# [resty](https://pkg.go.dev/github.com/go-resty/resty/v2)

## AddRetryCondition

```go
AddRetryCondition(func(r *resty.Response, err error) bool {
    // 保留原有的重试，不然会被覆盖
    if err != nil {
        return true
    }

    // 添加重试逻辑
})
```
