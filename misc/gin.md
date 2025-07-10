# gin

## excel响应

```golang
c.Header("content-type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
// 借助github.com/go-the-way/exl将结构体数组写入响应
exl.WriteTo(c.Writer, fruits)
```
