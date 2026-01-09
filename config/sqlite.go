package config

type SQLite struct {
	Path string `mapstructure:"path" json:"path"` // 数据库文件路径
}
