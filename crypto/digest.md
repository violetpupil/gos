# digest

除了使用16进制编码，还可以用base64编码

```golang
// md5
func main() {
 bytes := md5.Sum([]byte(s))
 digest := hex.EncodeToString(bytes[:])
 fmt.Println(digest)
}

// hmac
func main() {
 mac := hmac.New(sha1.New, []byte(key))
 mac.Write([]byte(s))
 bytes := mac.Sum(nil)
 digest := hex.EncodeToString(bytes)
 fmt.Println(digest)
}
```
