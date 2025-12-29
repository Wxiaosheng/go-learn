package db

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

/* 获取文章信息 */
func GetPostById(id int) (*Post, error) {
	var post Post
	result := sqlDB.First(&post, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}

/* 创建新文章 */
func CreatePost(post *Post) error {
	result := sqlDB.Create(post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/* 更新文章 */
func UpdatedPost(post *Post) (*Post, error) {
	result := sqlDB.Save(post)
	if result.Error != nil {
		return nil, result.Error
	}
	return post, nil
}
