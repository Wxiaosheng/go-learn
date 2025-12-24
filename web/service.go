package web

import (
	"github.com/gin-gonic/gin"
)

/*
go web 服务相关模块

1. 路由
*/
func InitService() {
	service := gin.Default()

	InitUserRoutes(service)

	service.Run(":8080")
}
