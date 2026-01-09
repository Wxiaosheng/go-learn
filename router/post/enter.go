package post

import api "go-learn/api/v1"

type RouterGroup struct {
	PostRouter
}

var (
	postApi = api.ApiGroupApp.PostApi
)
