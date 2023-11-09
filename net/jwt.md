# [jwt](https://jwt.io/)

[教程](https://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html)

服务端将数据存在客户端，通过校验签名防止数据篡改

## JWT 的数据结构

Header.Payload.Signature

Signature HmacSha256(base64UrlEncode(header)+"."+base64UrlEncode(payload), secret)

### Header 元数据

```json
{
  // 签名算法 HMAC SHA256
  "alg": "HS256",
  // token 类型 JWT
  "typ": "JWT"
}
```

### Payload 实际数据

通用字段，可另外包含自定义字段

`exp` 过期时间

## 前端处理

储存在 localStorage

请求头 Authorization: Bearer token
