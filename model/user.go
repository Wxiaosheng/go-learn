package model

// 存储用户信息，包括 id 、 username 、 password 、 email 等字段
type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"size:64;not null;uniqueIndex"`
	Password string
	Email    string `gorm:"size:255;not null;uniqueIndex"`
}
