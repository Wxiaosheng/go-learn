package global

import (
	"go-learn/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	BLOGS_VP *viper.Viper // Viper 全局变量

	BLOGS_LOG *zap.Logger // Zap 全局变量

	BLOGS_CONFIG config.BlogConfig // 配置文件全局变量
)
