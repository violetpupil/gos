# [csv](https://pkg.go.dev/encoding/csv)

字段可以用""包裹，""内可以换行，也可以使用""表示单个"

```golang
// ReadAll 读取csv文件
func ReadAll(name string) (records [][]string, err error) {
 f, err := os.Open(name)
 if err != nil {
  return nil, err
 }
 defer f.Close()

 reader := csv.NewReader(f)
 records, err = reader.ReadAll()
 return
}
```
