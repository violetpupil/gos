# [lo](https://pkg.go.dev/github.com/samber/lo)

```golang
// 去重
lo.Uniq(slice)
// 是否包含元素
lo.Contains(slice, item)
// 反转
lo.Reverse(slice)
// 切片转map
lo.SliceToMap(slice, transform)
// 切片用map分组
lo.GroupBy(slice, iteratee)
// 求和
lo.Sum(slice)
// 创建另一个类型切片
lo.Map(slice, iteratee)
// 创建另一个类型切片，可以过滤
lo.FilterMap(slice, f)
// 求交集
lo.Intersect(s1, s2)
// 求差集
lo.Difference(s1, s2)
```