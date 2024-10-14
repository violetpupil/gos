# oss

文件夹以/结尾

```golang
// 删除对象，不存在不会报错
// 只能删除空文件夹
b.DeleteObject("test.txt")
// 上传文件，key不能以斜杆开头
b.PutObject(key, io.Reader(resp.Body))
// 上传本地文件
b.PutObjectFromFile(objectKey, filePath)
```
