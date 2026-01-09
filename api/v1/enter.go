package v1

import "go-learn/service"

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	UserApi UserApi
	PostApi PostApi
}

var (
	userService = service.ServiceGroupApp.BasicService.UserService
	postService = service.ServiceGroupApp.PostService
)
