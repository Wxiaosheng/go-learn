package post

import (
	"go-learn/global"
	"go-learn/model"
)

type PostService struct{}

/* 根据文章ID判断当前用户是否是所有者 */
func (ps *PostService) IsPostOwner(postId int, userId int) bool {
	var post model.Post
	result := global.BLOGS_DB.First(&post, postId)
	if result.Error != nil {
		return false
	}
	return post.UserID == userId
}

/* 获取文章 */
func (ps *PostService) GetPostById(post *model.Post) error {
	result := global.BLOGS_DB.First(&post, post.ID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/* 创建文章 */
func (ps *PostService) CreatePost(post *model.Post) error {
	result := global.BLOGS_DB.Create(post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/* 更新文章 */
func (ps *PostService) UpdatePost(post *model.Post) error {
	result := global.BLOGS_DB.Save(post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/* 删除文章 */
func (ps *PostService) DeletePost() {

}
