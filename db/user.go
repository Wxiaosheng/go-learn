package db

// 存储用户信息，包括 id 、 username 、 password 、 email 等字段
type User struct {
	ID       int    `GORM:"primaryKey"`
	Username string `GORM:"uniqueIndex"`
	Password string
	Email    string `GORM:"uniqueIndex"`
}

/*
获取用户信息
*/
func GetUserInfo(id int) *User {
	var user User
	result := sqlDB.First(&user, id)
	if result.Error != nil {
		return nil
	}
	return &user
}
