package web

import (
	"go-learn/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(service *gin.Engine) {
	// 获取用户信息
	service.GET("/user/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{"error": "请传递有效的用户 ID"})
			return
		}
		user := db.GetUserInfo(id)
		if user == nil {
			ctx.JSON(404, gin.H{"error": "用户不存在"})
			return
		}
		ctx.JSON(200, user)
	})

	// 用户注册
	service.POST("/user/sign", func(ctx *gin.Context) {

	})

	// 用户登录
	service.POST("/user/login", func(ctx *gin.Context) {

	})
}
