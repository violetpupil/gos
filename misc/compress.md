# compress

## [gzip](https://pkg.go.dev/compress/gzip)

## [brotli](https://pkg.go.dev/github.com/andybalholm/brotli)

```golang
// 压缩
func BrotliCompress(p []byte) (string, error) {
 var buffer bytes.Buffer

 options := brotli.WriterOptions{
  Quality: 3,
 }
 writer := brotli.NewWriterOptions(&buffer, options)

 _, err := writer.Write(p)
 if err != nil {
  return "", errors.Wrap(err, "write")
 }

 err = writer.Close()
 if err != nil {
  return "", errors.Wrap(err, "close")
 }

 result := base64.StdEncoding.EncodeToString(buffer.Bytes())
 return result, nil
}

// 解压缩
func BrotliDecompress(p string) ([]byte, error) {
 data, err := base64.StdEncoding.DecodeString(p)
 if err != nil {
  return nil, errors.Wrap(err, "DecodeString")
 }

 reader := brotli.NewReader(bytes.NewReader(data))

 var buffer bytes.Buffer
 _, err = io.Copy(&buffer, reader)
 if err != nil {
  return nil, errors.Wrap(err, "Copy")
 }

 result := buffer.Bytes()
 return result, nil
}
```
