package router

import "go-learn/router/basic"

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	Basic basic.RouterGroup
}
