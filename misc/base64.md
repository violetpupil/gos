# [base64](https://pkg.go.dev/encoding/base64)

每6个bit位编码一个字节，三个字节会变成四个字节，增大三分之一

如果最后剩两个字节，也就是16位，则补充两个0位，编码成三个字节，最后一个字节是=

如果最后剩一个字节，也就是8位，则补充四个0位，编码成两个字节，最后两个字节是==

```golang
base64.StdEncoding.EncodeToString
base64.StdEncoding.DecodeString
base64.URLEncoding.EncodeToString
base64.URLEncoding.DecodeString
```

## StdEncoding

```text
0-25 -> A-Z
26-51 -> a-z
52-61 -> 0-9
62 -> +
63 -> /
```

## URLEncoding

```text
0-25 -> A-Z
26-51 -> a-z
52-61 -> 0-9
62 -> -
63 -> _
```
