package basic

import api "go-learn/api/v1"

type RouterGroup struct {
	UserRouter
}

var (
	userApi = api.ApiGroupApp.UserApi
)
