package db

import (
	"errors"
)

// 存储用户信息，包括 id 、 username 、 password 、 email 等字段
type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Password string
	Email    string `gorm:"uniqueIndex"`
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

/*
用户注册
*/
func UserSign(user User) error {
	// 注册用户
	err := sqlDB.Create(&user).Error

	if err != nil {
		if IsDuplicatedKeyError(err) {
			return errors.New("用户名或邮箱已存在")
		}
		return err
	}

	return nil
}
