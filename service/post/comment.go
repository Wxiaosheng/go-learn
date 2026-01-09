package post

import (
	"go-learn/global"
	"go-learn/model"

	systemRes "go-learn/model/response"
)

/** 创建评论 */
func (ps *PostService) CreateComment(comment *model.Comment) error {
	result := global.BLOGS_DB.Create(comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/** 获取文章的评论列表 */
func (ps *PostService) GetCommentsByPostId(postId int) ([]systemRes.CommentRes, error) {
	var comments []systemRes.CommentRes
	result := global.BLOGS_DB.Where("post_id = ?", postId).Order("created_at DESC").Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}

	// 构建 嵌套结构 的数据
	commentMap := make(map[int]*systemRes.CommentRes)

	for i := range comments {
		comment := &comments[i]
		commentMap[comment.ID] = comment
		comment.Replies = []systemRes.CommentRes{} // 初始化Replies（避免 nil）
	}

	// 构建根评论列表
	rootComments := []systemRes.CommentRes{}

	for _, comment := range comments {
		if comment.ParentID == 0 { // 文章评论
			rootComments = append(rootComments, comment)
		} else { // 子评论
			if parent, exists := commentMap[comment.ParentID]; exists { // 存在父评论（可能出现不存在父评论的情况）
				parent.Replies = append(parent.Replies, comment)
			}
		}
	}

	return rootComments, nil
}
