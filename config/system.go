package config

type System struct {
	Port   int    `mapstructure:"port"`    // 端口号
	DbType string `mapstructure:"db-type"` // 数据库类型
}
