package basic

import (
	"errors"
	"go-learn/model"
	"go-learn/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (us *UserService) UserInfo(u *model.User) {
	u.GetUserInfoById()
}

/* 处理用户注册 */
func (us *UserService) UserSign(u *model.User) error {
	// 密码存储处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return u.UserSign()
}

/* 用户登陆 */
func (us *UserService) UserLogin(u *model.User) (string, error) {
	// 1、根据用户名获取用户信息
	user := u.GetUserInfoByName()
	if user == nil {
		return "", errors.New("用户不存在")
	}

	// 2、验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		return "", errors.New("密码错误")
	}

	// 3、jwt 生成
	tokenString, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", err
	}

	// 4、返回 token
	return tokenString, nil
}
