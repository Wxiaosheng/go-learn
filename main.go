package main

import (
	"fmt"
	"go-learn/global"
	"go-learn/initial"
)

func main() {
	/* 初始化服务基础设施 */

	initial.InitViper() // 初始化 Viper 配置
	initial.InitZap()   // 初始化 Zap 日志

	global.BLOGS_LOG.Info("配置文件和日志初始化完成")
	// 启动 web 服务
	fmt.Printf("启动服务")
}
