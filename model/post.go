package model

// 存储博客文章信息，包括 id 、 title 、 content 、 user_id （关联 users 表的 id ）、 created_at 、 updated_at 等字段
type Post struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Title     string `json:"title" gorm:"not null"`
	Content   string `json:"content"`
	UserID    int    `json:"user_id" gorm:"not null"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at,omitempty"`
}

// 存储评论信息，包括 id 、 post_id （关联 posts 表的 id ）、 user_id （关联 users 表的 id ）、 content 、 created_at 等字段
type Comment struct {
	ID        int    `json:"id" gorm:"primaryKey; comment: '评论ID'"`
	Content   string `json:"content" gorm:"not null; size:2000; comment: '评论内容'"`
	PostID    int    `json:"post_id" gorm:"not null; index; comment: '关联的文章ID（外键）'"`
	UserID    int    `json:"user_id" gorm:"not null; index; comment: '评论者用户ID（外键）'"`
	ParentID  int    `json:"parent_id,omitempty" gorm:"default:0; comment: '回复的评论ID（0表示顶级评论，非0表示回复某条评论）'"`
	CreatedAt int64  `json:"created_at"`
	DeletedAt int64  `json:"deleted_at,omitempty"`
}
