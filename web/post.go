package web

import (
	"errors"
	"go-learn/db"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PostData struct {
	ID      int    `json:"id,omitempty"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Create  int64  `json:"created_at,omitempty"`
	Update  int64  `json:"updated_at,omitempty"`
}

// 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容
func InitPostRoutes(service *gin.Engine) {
	// 实现文章的更新功能，只有文章的作者才能更新自己的文章
	// 有 ID 是更新文章，否则是创建文章
	service.POST("/post", func(ctx *gin.Context) {
		var postData PostData
		if err := ctx.ShouldBindJSON(&postData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "JSON 解析失败"})
			return
		}
		userId, _ := ctx.Get("userId")

		if postData.ID == 0 {
			// 创建文章
			db.CreatePost(&db.Post{
				Title:     postData.Title,
				Content:   postData.Content,
				CreatedAt: time.Now().UnixMilli(),
				UpdatedAt: time.Now().UnixMilli(),
				UserID:    userId.(int),
			})
			ctx.JSON(http.StatusOK, gin.H{"msg": "文章创建成功"})
		} else {
			// 更新文章
			// 仅作者才能更新
			post, err := db.GetPostById(postData.ID)
			if err != nil {
				ctx.JSON(http.StatusNotFound, err.Error())
				return
			}
			if post.UserID != userId.(int) {
				ctx.JSON(http.StatusForbidden, gin.H{"msg": "无权更新此文章"})
				return
			}
			db.UpdatedPost(&db.Post{
				ID:        postData.ID,
				Title:     postData.Title,
				Content:   postData.Content,
				UpdatedAt: time.Now().UnixMilli(),
			})
			ctx.JSON(http.StatusOK, gin.H{"msg": "文章更新成功"})
		}

	})

	// 实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
	service.POST("/getPost", func(ctx *gin.Context) {
		postId, err := getPostId(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, err.Error())
			return
		}
		post, err := db.GetPostById(postId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"msg": "文章不存在"})
			return
		}
		ctx.JSON(http.StatusOK, post)
	})

	// 实现文章的删除功能，只有文章的作者才能删除自己的文章。
	service.POST("/post/delete", func(ctx *gin.Context) {
		postId, err := getPostId(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, err.Error())
			return
		}
		userId, _ := ctx.Get("userId")
		post, err := db.GetPostById(postId)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}

		if post.UserID != userId.(int) {
			ctx.JSON(http.StatusForbidden, gin.H{"msg": "无权删除此文章"})
			return
		}

		post.DeletedAt = time.Now().UnixMilli()
		_, err = db.UpdatedPost(post)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "删除文章失败"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"msg": "文章删除成功"})
	})
}

type postId struct {
	ID int `json:"id" binding:"required"`
}

func getPostId(ctx *gin.Context) (int, error) {
	var p postId
	if err := ctx.ShouldBindJSON(&p); err != nil {
		return 0, errors.New("缺少文章 ID")
	}
	return p.ID, nil
}
