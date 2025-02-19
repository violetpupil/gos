# [deepcopy](https://pkg.go.dev/github.com/mohae/deepcopy)

```golang
src := map[string]int{}

// 浅拷贝，只复制了地址，赋值会相互影响
dst := src

// 深拷贝
dstCp := deepcopy.Copy(src)
dst := dstCp.(map[string]int)
```
