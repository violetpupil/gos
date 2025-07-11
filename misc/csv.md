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

## [gocsv](https://pkg.go.dev/github.com/gocarina/gocsv)

```golang
// UnmarshalFile 读取csv文件
// 可以使用结构体标签csv，与文件第一行对应
func UnmarshalFile(name string, out any) error {
 f, err := os.Open(name)
 if err != nil {
  return err
 }
 defer f.Close()

 err = gocsv.UnmarshalFile(f, out)
 return err
}

// MarshalFile 写csv文件
func MarshalFile(name string, in any) error {
 f, err := os.Create(name)
 if err != nil {
  return err
 }
 defer f.Close()

 err = gocsv.MarshalFile(in, f)
 return err
}

// 改变全局reader的字段分隔符，默认是逗号,
gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
  r := csv.NewReader(in)
  r.Comma = comma
  return r
})
```
