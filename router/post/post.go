package post

import "github.com/gin-gonic/gin"

type PostRouter struct{}

/* 初始化 post 服务路由 */
func (pr *PostRouter) InitPostRouter(Router *gin.RouterGroup) {
	postRouter := Router.Group("/post")

	{
		postRouter.POST("/", postApi.UpdatePost)
		postRouter.POST("/getPost", postApi.GetPostInfo)
		postRouter.POST("/delete", postApi.DeletePost)
	}
}
