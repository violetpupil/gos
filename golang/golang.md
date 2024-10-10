# golang

## generics

```golang
// 函数声明
func F[V any](v V)
// 多个类型参数
func F[K any, V any](k K, v V)
// 类型约束多个类型
func F[V constraints.Float | constraints.Integer](v V)
// 约束原类型和重定义类型(type Int int)
func F[V ~int](v V)

// 类型约束声明
type Number interface {
    constraints.Float | constraints.Integer
}
```