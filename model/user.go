package model

import (
	"errors"
	"go-learn/global"
)

// 存储用户信息，包括 id 、 username 、 password 、 email 等字段
type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"size:64;not null;uniqueIndex"`
	Password string
	Email    string `gorm:"size:255;not null;uniqueIndex"`
}

/* 获取用户信息 */
func (user *User) GetUserInfoById() *User {
	result := global.BLOGS_DB.First(&user, user.ID)
	if result.Error != nil {
		return nil
	}
	return user
}

/* 获取用户信息 */
func (user *User) GetUserInfoByName() *User {
	if result := global.BLOGS_DB.Where("username = ?", user.Username).First(&user); result.Error != nil {
		return nil
	}
	return user
}

/* 用户注册 */
func (user *User) UserSign() error {
	// 注册用户
	err := global.BLOGS_DB.Create(&user).Error

	if err != nil {
		if IsDuplicatedKeyError(err) {
			return errors.New("用户名或邮箱已存在")
		}
		return err
	}

	return nil
}
