package web

import (
	"go-learn/middleware"

	"github.com/gin-gonic/gin"
)

/*
go web 服务相关模块

1. 路由
*/
func InitService() {
	service := gin.Default()

	// 全局使用 JWT 鉴权中间件
	service.Use(middleware.JwtAuthMiddleware())

	InitUserRoutes(service)
	InitPostRoutes(service)

	service.Run(":8080")
}
