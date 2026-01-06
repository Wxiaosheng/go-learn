package v1

import (
	"go-learn/model"
	systemReq "go-learn/model/request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

/* 获取用户信息 */
func (up *UserApi) GetUserInfo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "请传递有效的用户 ID"})
		return
	}

	user := &model.User{ID: id}
	userService.UserInfo(user)

	if user == nil {
		ctx.JSON(404, gin.H{"error": "用户不存在"})
		return
	}
	ctx.JSON(200, user)
}

/* 用户注册 */
func (up *UserApi) UserSign(ctx *gin.Context) {
	var userData systemReq.UserRes
	if err := ctx.ShouldBindJSON(&userData); err != nil {
		ctx.JSON(400, gin.H{"msg": "JSON 解析失败"})
		return
	}

	err := userService.UserSign(&model.User{
		Username: userData.UserName,
		Password: userData.Password,
		Email:    userData.Email,
	})
	if err != nil {

		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"msg": "success"})
}

/* 用户登录 */
func (up *UserApi) UserLogin(ctx *gin.Context) {
	var user systemReq.UserRes
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "JSON 解析失败"})
		return
	}

	token, err := userService.UserLogin(&model.User{
		Username: user.UserName,
		Password: user.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
		return
	}

	ctx.Header("Authorization", token)
	ctx.JSON(http.StatusOK, gin.H{"msg": "登录成功"})
}
