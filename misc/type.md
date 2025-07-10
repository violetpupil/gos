# type

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

## map

map没有初始化，可以取值，但不能添加元素

引用类型，赋值是浅拷贝，深拷贝可以用deepcopy库
