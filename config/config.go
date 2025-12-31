package config

/* 配置相关结构体定义 */
type BlogConfig struct {
	System System `mapstructure:"system"`

	Zap Zap `mapstructure:"zap"`

	Mysql MySQL `mapstructure:"mysql"`

	Sqlite SQLite `mapstructure:"sqlite"`
}
