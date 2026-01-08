package router

import (
	"go-learn/router/basic"
	"go-learn/router/post"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	Basic basic.RouterGroup
	Post  post.RouterGroup
}
