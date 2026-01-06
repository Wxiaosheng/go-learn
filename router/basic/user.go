package basic

import "github.com/gin-gonic/gin"

type UserRouter struct{}

/* 初始化 User 模块路由 */
func (ur *UserRouter) InitUserRouters(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	{
		userRouter.GET("/:id", userApi.GetUserInfo)
		userRouter.POST("/sign", userApi.UserSign)
		userRouter.POST("/login", userApi.UserLogin)
	}
}
