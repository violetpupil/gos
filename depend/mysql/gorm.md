# gorm

## 分页

```golang
// 有一个小于0则查所有
if req.Page == 0 {
 req.Page = 1
}
if req.PageSize == 0 {
 req.PageSize = 10
}

if req.Page > 0 && req.PageSize > 0 {
 db.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize)
}
```
