package resty

import "github.com/go-resty/resty/v2"

// ReqHook 设置请求信息
// 查询字符串，键区分大小写
// SetQueryParam 单个参数，同一个键会覆盖
// SetQueryParams 多个参数
// SetQueryParamsFromValues 多个参数，参数可能有多个值
// SetQueryString 直接设置字符串
//
// 请求头
// 键不区分大小写，规范格式 Content-Type
// 多个值格式为 "text/html, application/xml;q=0.9"
// SetHeader 单个参数，同一个键会覆盖
// SetHeaderMultiValues 多个参数，参数可能有多个值
// SetHeaders 多个参数，同一个键会覆盖
//
// 请求体 SetBody
// 支持类型`string`, `[]byte`, `struct`, `map`, `slice`, `io.Reader`.
// 支持指针，struct、map、slice类型会发送json请求体
//
// 认证请求头 `Authorization: <scheme> <token>`
// https://www.iana.org/assignments/http-authschemes/http-authschemes.xhtml
// SetAuthScheme 认证方案，默认为Bearer - OAuth 2.0
// SetAuthToken 认证令牌
// SetBasicAuth basic认证 "Basic <base64-encoded-value>"
//
// 响应体自动解码，与Response.Result配合使用
// 响应码为200~299，响应体类型为json或xml
// SetResult 指定响应体结构
//
// EnableTrace 是否追踪请求信息
type ReqHook = func(*resty.Request) *resty.Request
