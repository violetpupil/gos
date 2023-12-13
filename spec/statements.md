# Statements

## for range

```golang
s := []int{0, 1, 2}
// 声明 int 类型变量 v
// 每次循环进行赋值
// 变量 v 的地址不变
// 作用域在 for 语句中
for _, v := range s {
 fmt.Println(&v)
}

for _, v := range s {
 // 每次循环，使用新的变量
 v := v
 fmt.Println(&v)
}
```
