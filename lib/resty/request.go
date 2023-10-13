package resty

import "github.com/go-resty/resty/v2"

// 请求
// Post() post方法
// 日志
// 请求前 路径 请求体 自定义日志
// 请求后 路径 请求体 自定义日志 响应码+描述 响应体 耗时
type Request = resty.Request

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
// cookie
// SetCookie 单个cookie
// SetCookies 多个cookie
//
// 请求体 SetBody
// 支持类型`string`, `[]byte`, `struct`, `map`, `slice`, `io.Reader`，支持指针
// struct、map、slice类型自动json编码，并设置 Content-Type 为 application/json
// string类型 Content-Type 为 "text/plain; charset=utf-8"
// []byte类型 Content-Type 根据内容检测，保底为 "text/plain; charset=utf-8"
//
// 表单
// 只传键值 Content-Type 为 application/x-www-form-urlencoded
// 上传文件 Content-Type 为 multipart/form-data
// SetFile 上传单文件
// SetFileReader 上传单文件，使用io.Reader
// SetFiles 上传多文件
// SetFormData 设置表单键值对
// SetFormDataFromValues 设置表单键值对，支持多值
//
// 认证请求头 `Authorization: <scheme> <token>`
// https://www.iana.org/assignments/http-authschemes/http-authschemes.xhtml
// SetAuthScheme 认证方案，默认为Bearer - OAuth 2.0
// SetAuthToken 认证令牌
// SetBasicAuth basic认证 "Basic <base64-encoded-value>"
//
// 响应体自动解码，当类型为json或xml时
// SetResult 指定响应体结构，响应码200~299
// 当指定结构非指针时，与Response.Result配合使用
// 当指定结构为指针时，与Response.Result配合使用，或直接使用传入的指针
// SetError 指定响应体结构，响应码>399
// 当指定结构非指针时，与Response.Error配合使用
// 当指定结构为指针时，与Response.Error配合使用，或直接使用传入的指针
//
// EnableTrace 是否追踪请求信息
// SetOutput 指定下载路径
type ReqHook = func(*resty.Request)
