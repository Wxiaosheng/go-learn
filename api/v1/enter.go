package v1

import "go-learn/service"

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	UserApi UserApi
}

var (
	userService = service.ServiceGroupApp.BasicService.UserService
)
