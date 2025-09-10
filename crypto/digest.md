# digest

```golang
// md5
func main() {
 bytes := md5.Sum([]byte(s))
 digest := hex.EncodeToString(bytes[:])
 fmt.Println(digest)
}
```
