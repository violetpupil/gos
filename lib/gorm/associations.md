# associations

## [Belongs To](https://gorm.io/docs/belongs_to.html)

```golang
// 默认使用外键 类型名+主键名
type User struct {
 CompanyID int
 Company   Company
}

type Company struct {
 ID int
}
```

```golang
// 自定义外键
type User struct {
 CompanyRefer int
 Company      Company `gorm:"foreignKey:CompanyRefer"`
}
```

```golang
// 自定义外键参考字段
type User struct {
 CompanyID int
 Company   Company `gorm:"references:Code"`
}
```

## [Has One](https://gorm.io/docs/has_one.html)

```golang
type User struct {
 ID         uint
 CreditCard CreditCard
}

type CreditCard struct {
 UserID uint
}
```

## [Has Many](https://gorm.io/docs/has_many.html)

```golang
type User struct {
 ID          uint
 CreditCards []CreditCard
}
```

## [Many To Many](https://gorm.io/docs/many_to_many.html)

Many to Many add a join table between two models.

the join table owns the foreign key which references two models

`joinForeignKey` `joinReferences` 标签 指定中间表外键

```golang
// 指定中间表
type User struct {
 Languages []Language `gorm:"many2many:user_languages"`
}

type Language struct {
 Users []User `gorm:"many2many:user_languages"`
}
```
