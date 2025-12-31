package config

/* 配置相关结构体定义 */
type BlogConfig struct {
	Mysql MySQL `mapstructure:"mysql"`

	Sqlite SQLite `mapstructure:"sqlite"`
}
