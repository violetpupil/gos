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

## 删除

```golang
db.Where("id = ?", 1).Delete(&User{})

// 软删除
type User struct {
  DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
```
