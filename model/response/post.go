package response

import "go-learn/model"

type CommentRes struct {
	model.Comment

	Replies []CommentRes `json:"replies" gorm:"-"`
}

// 显式指定 CommentRes 结构体关联的数据库表名
func (cRes *CommentRes) TableName() string {
	return "comments"
}
