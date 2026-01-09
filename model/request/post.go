package request

type PostId struct {
	ID int `json:"id" binding:"required"`
}

type PostRes struct {
	ID      int    `json:"id,omitempty"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Create  int64  `json:"createdAt,omitempty"`
	Update  int64  `json:"updatedAt,omitempty"`
}

// 创建评论请求结构体
type CommentRes struct {
	ID       int    `json:"id,omitempty"`
	PostID   int    `json:"postId" binding:"required"`
	ParentID int    `json:"parentId,omitempty"`
	Content  string `json:"content" binding:"required"`
}

// 获取评论列表请求结构体
type GetCommentsReq struct {
	PostID int `json:"postId" binding:"required"`
}
