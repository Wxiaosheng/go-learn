package initial

import (
	"go-learn/global"

	"go.uber.org/zap"
)

func InitZap() {
	log, _ := zap.NewProduction()
	defer log.Sync() // 刷新缓冲区

	global.BLOGS_LOG = log
}
