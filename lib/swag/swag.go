// https://github.com/swaggo/swag
// go install github.com/swaggo/swag/cmd/swag@latest
// swag init
// swaggertype 结构体标签
// swaggerignore 结构体标签
//
// General API Info
// @title 必须
// @version 必须
//
// API Operation
// @router 路由信息 必要字段
// @summary A short summary of what the operation does.
// @description A verbose explanation of the operation behavior.
// @tags 逗号分隔的标签列表
// @param 参数名 参数类型 数据类型 是否必须true "comment"
// @response 200 {object} 数据类型 "comment"
// 示例
// func 接口作用
// @router /echo [post]
// @summary 接口作用
// @description 接口作用
// @tags 标签
// @param req body struct true "req"
// @response 200 {object} struct
//
// Param Type
// query path header body formData
//
// Data Type
// string int uint uint32 uint64 float32 bool 自定义类型
// file 上传文件
// 指定字段类型 Response{d1=string,d2=[]string}
package swag

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// GinSwagger gin swagger文档服务
// 访问路径 /swagger/index.html
// 必须导入生成的docs包
func GinSwagger(e *gin.Engine) {
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
