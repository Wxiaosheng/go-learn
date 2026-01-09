package request

type PostId struct {
	ID int `json:"id" binding:"required"`
}

type PostRes struct {
	ID      int    `json:"id,omitempty"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Create  int64  `json:"created_at,omitempty"`
	Update  int64  `json:"updated_at,omitempty"`
}
