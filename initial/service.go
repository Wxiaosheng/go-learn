package initial

import (
	"fmt"
	"go-learn/global"
	"go-learn/middleware"
	"go-learn/router"

	"github.com/gin-gonic/gin"
)

/* 初始化服务 */
func InitService() {
	service := gin.New()

	service.Use(middleware.JwtAuthMiddleware())

	basicRouter := router.RouterGroupApp.Basic
	postRouter := router.RouterGroupApp.Post

	PublicGroup := service.Group(global.BLOGS_CONFIG.System.RouterPrefix)
	// PrivateGroup := service.Group(global.BLOGS_CONFIG.System.RouterPrefix)

	{ // basic 模块路由
		basicRouter.InitUserRouters(PublicGroup)
	}

	{ // post 模块路由
		postRouter.InitPostRouter(PublicGroup)
	}

	port := fmt.Sprintf(":%d", global.BLOGS_CONFIG.System.Port)
	service.Run(port)
}
