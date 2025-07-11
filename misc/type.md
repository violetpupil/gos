# type

[lo](https://pkg.go.dev/github.com/samber/lo)

## [slices](https://pkg.go.dev/slices)

```golang
// 原地升序
slices.Sort(slice)
// 排序任意切片
// cmp函数接收两个切片元素作为参数a和b，返回int
// 升序 a=b返回0，a>b返回1，a<b返回-1
// 降序 a=b返回0，a>b返回-1，a<b返回1
slices.SortFunc(slice, cmp)
```

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

## map

map没有初始化，可以取值，但不能添加元素

引用类型，赋值是浅拷贝，深拷贝可以用deepcopy库

## strings

```golang
// 切分字符串
// sep不为空，且str不包含sep，则返回[str]
strings.Split(str, sep)
// 判断字符串相等，大小写不敏感
strings.EqualFold(s1, s2)
```
