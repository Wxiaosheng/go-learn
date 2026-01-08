package v1

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go-learn/model"
	systemReq "go-learn/model/request"
)

type PostApi struct{}

// 有 ID 是更新文章，否则是创建文章
func (pa *PostApi) UpdatePost(ctx *gin.Context) {
	var postData systemReq.PostRes
	if err := ctx.ShouldBindJSON(&postData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "JSON 解析失败"})
		return
	}
	userId, _ := ctx.Get("userId")

	post := &model.Post{
		ID:        postData.ID,
		Title:     postData.Title,
		Content:   postData.Content,
		CreatedAt: time.Now().UnixMilli(),
		UpdatedAt: time.Now().UnixMilli(),
		UserID:    userId.(int),
	}

	if postData.ID == 0 {
		// 创建文章
		postService.CreatePost(post)
		ctx.JSON(http.StatusOK, gin.H{"msg": "文章创建成功"})
	} else {
		// 更新文章
		// 仅作者才能更新
		err := postService.GetPostById(post)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
		if post.UserID != userId.(int) {
			ctx.JSON(http.StatusForbidden, gin.H{"msg": "无权更新此文章"})
			return
		}

		postService.UpdatePost(post)
		ctx.JSON(http.StatusOK, gin.H{"msg": "文章更新成功"})
	}

}

// 实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
func (pa *PostApi) GetPostInfo(ctx *gin.Context) {
	postId, err := getPostId(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}

	post := &model.Post{ID: postId}

	if err := postService.GetPostById(post); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "文章不存在"})
		return
	}
	ctx.JSON(http.StatusOK, post)
}

// 实现文章的删除功能，只有文章的作者才能删除自己的文章。
func (pa *PostApi) DeletePost(ctx *gin.Context) {
	postId, err := getPostId(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}
	userId, _ := ctx.Get("userId")

	post := &model.Post{ID: postId}

	if err := postService.GetPostById(post); err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	if post.UserID != userId.(int) {
		ctx.JSON(http.StatusForbidden, gin.H{"msg": "无权删除此文章"})
		return
	}

	post.DeletedAt = time.Now().UnixMilli()

	if err := postService.UpdatePost(post); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "删除文章失败"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "文章删除成功"})
}

func getPostId(ctx *gin.Context) (int, error) {
	var p systemReq.PostId
	if err := ctx.ShouldBindJSON(&p); err != nil {
		return 0, errors.New("缺少文章 ID")
	}
	return p.ID, nil
}
