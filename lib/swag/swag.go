// go install github.com/swaggo/swag/cmd/swag@latest
// swag init
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
