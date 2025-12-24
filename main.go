package main

import (
	"go-learn/db"
	"go-learn/web"
)

func main() {

	// 任务一
	// homework.ExectTask01()

	// 任务二
	// homework.ExectTask02()

	// 任务三
	// homework.Task03()
	// homework.Test03_2()

	// 初始化表结构
	db.InitDB()

	// 初始化 web 服务
	web.InitService()
}
