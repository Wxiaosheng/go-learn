package main

import (
	"fmt"
	"go-learn/global"
	"go-learn/initial"
)

func main() {
	/* 初始化服务基础设施 */

	initial.InitViper() // 初始化 Viper 配置

	fmt.Printf("%+v", global.BLOGS_CONFIG)

	// 启动 web 服务
	fmt.Printf("启动服务")
}
