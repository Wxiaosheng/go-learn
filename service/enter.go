package service

import (
	"go-learn/service/basic"
	"go-learn/service/post"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	BasicService basic.BasicService
	PostService  post.PostServiceGroup
}
