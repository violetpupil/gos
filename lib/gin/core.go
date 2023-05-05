package gin

import "github.com/gin-gonic/gin"

type (
	// Context 请求上下文
	// MustBindWith() 解析参数，解析失败会返回错误并响应400
	// ShouldBindWith() 解析参数，解析失败只返回错误
	// BindHeader() 解析请求头，结构体标签header
	// BindJSON() 解析json请求体
	// BindQuery() 解析查询字符串
	// BindUri() 解析路径参数，结构体标签uri
	// Bind() 自动选择参数类型
	// Bind()可以同时解析查询字符串和url编码表单
	// 结构体标签都是form，表单优先级高
	// 结构体标签 binding:"required" 校验字段不能为空
	//
	// JSON() 生成json响应体，并设置Content-Type
	Context = gin.Context
	// Engine 服务引擎
	// Use() 配置全局中间件
	Engine = gin.Engine
	// RouterGroup 路由组
	// Handle() 注册请求处理，路径可不以/开头
	// GET() get处理
	// POST() post处理
	//
	// Group() 创建子路由组，路径可不以/结尾
	// Use() 配置中间件
	RouterGroup = gin.RouterGroup
)