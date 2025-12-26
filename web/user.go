package web

import (
	"errors"
	"go-learn/db"
	"go-learn/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"omitempty,email"`
}

func InitUserRoutes(service *gin.Engine) {
	// 获取用户信息
	service.GET("/user/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, gin.H{"error": "请传递有效的用户 ID"})
			return
		}
		user := db.GetUserInfoById(id)
		if user == nil {
			ctx.JSON(404, gin.H{"error": "用户不存在"})
			return
		}
		ctx.JSON(200, user)
	})

	// 用户注册
	service.POST("/user/sign", func(ctx *gin.Context) {
		var userData UserData
		if err := ctx.ShouldBindJSON(&userData); err != nil {
			ctx.JSON(400, gin.H{"msg": "JSON 解析失败"})
			return
		}

		err := handleSign(&userData)
		if err != nil {

			ctx.JSON(400, gin.H{"msg": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"msg": "success"})
	})

	// 用户登录
	service.POST("/user/login", func(ctx *gin.Context) {
		var user UserData
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "JSON 解析失败"})
			return
		}

		token, err := handleLogin(&user)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
			return
		}

		ctx.Header("Authorization", token)
		ctx.JSON(http.StatusOK, gin.H{"msg": "登录成功"})
	})
}

/* 处理用户注册 */
func handleSign(u *UserData) error {
	// 密码存储处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return db.UserSign(db.User{
		Username: u.UserName,
		Password: string(hashedPassword),
		Email:    u.Email,
	})
}

/* 用户登陆 */
func handleLogin(u *UserData) (string, error) {
	// 1、根据用户名获取用户信息
	user := db.GetUserInfoByName(u.UserName)
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
