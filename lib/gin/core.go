package gin

import "github.com/gin-gonic/gin"

type (
	// ResponseWriter 响应构造器
	// Status() HTTP response status code
	ResponseWriter = gin.ResponseWriter
	// Context 请求上下文
	// Request http请求
	// Writer 响应构造器
	//
	// MustBindWith() 解析参数，解析失败会返回错误并响应400
	// ShouldBindWith() 解析参数，解析失败只返回错误
	// BindHeader() 解析请求头，结构体标签header
	// BindJSON() ShouldBindJSON() 解析json请求体
	// BindQuery() 只解析查询字符串，结构体标签form
	// BindUri() 解析路径参数，结构体标签uri
	// Bind() 自动选择参数类型
	// Bind()可以同时解析查询字符串和url编码表单
	// 结构体标签都是form，表单优先级高
	// 结构体标签 binding:"required" 校验字段不能为空
	//
	// Abort() 放弃剩下的处理器
	// AbortWithStatus() 指定状态码
	// AbortWithStatusJSON() 指定状态码并响应json
	// Next() 嵌套执行之后的处理器
	// 不调用 Next() 剩下的处理器也会执行
	//
	// Header() 设置响应头，如果值为""，移除响应头
	// JSON() 生成json响应体，并设置Content-Type
	// Status() sets the HTTP response code.
	// ClientIP() 客户端ip
	// Cookie() 返回指定cookie值，找不到返回错误
	Context = gin.Context
	// Engine 服务引擎
	// Use() 配置全局中间件
	Engine = gin.Engine
	// RouterGroup 路由组
	// Handle() 注册请求处理，路径可不以/开头
	// Any() 所有请求方法
	// GET() get处理
	// POST() post处理
	//
	// Group() 创建子路由组，路径可不以/结尾
	// Use() 配置中间件
	RouterGroup = gin.RouterGroup
	// HandlerFunc 处理器函数
	HandlerFunc = gin.HandlerFunc
)
