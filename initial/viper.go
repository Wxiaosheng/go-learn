package initial

import (
	"go-learn/global"

	"github.com/spf13/viper"
)

/* Viper 配置初始化 */
func InitViper() {
	viper := viper.New()

	viper.SetConfigName("config") // 配置文件名（不带扩展名）
	viper.SetConfigType("yaml")   // 配置文件格式
	viper.AddConfigPath(".")      // 配置文件路径

	err := viper.ReadInConfig() // 读取配置文件
	if err != nil {
		panic("读取配置文件失败: " + err.Error())
	}

	if err := viper.Unmarshal(&global.BLOGS_CONFIG); err != nil { // 将配置文件内容映射到全局变量中
		panic("fatal error config file: " + err.Error())
	}

	// 将 viper 对象赋值到全局变量中
	global.BLOGS_VP = viper
}
