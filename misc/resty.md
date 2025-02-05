# [resty](https://pkg.go.dev/github.com/go-resty/resty/v2)

## AddRetryCondition

```go
// 保留原有的重试，不然会被覆盖
AddRetryAfterErrorCondition().
AddRetryCondition(func(r *resty.Response, err error) bool {})
```
