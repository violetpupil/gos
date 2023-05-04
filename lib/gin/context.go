// Context
//
// 结构体标签 binding:"required" 校验字段不能为空
// 使用Bind()或ShouldBind()同时解析查询字符串和url编码表单
// 如果字段重合，以表单为准
// 结构体标签form
//
// MustBindWith() 解析参数，解析失败会返回错误并响应400
// BindHeader() 解析请求头，结构体标签header
// BindJSON() 解析json请求体
// BindQuery() 解析查询字符串
// BindUri() 解析路径参数，结构体标签uri
// Bind() 自动选择参数类型
//
// ShouldBindWith() 解析参数，解析失败只返回错误
// ShouldBindHeader() 解析请求头，结构体标签header
// ShouldBindJSON() 解析json请求体
// ShouldBindQuery() 解析查询字符串
// ShouldBindUri() 解析路径参数，结构体标签uri
// ShouldBind() 自动选择参数类型
//
// JSON() 生成json响应体，并设置Content-Type
package gin
