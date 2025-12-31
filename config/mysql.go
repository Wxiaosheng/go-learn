package config

type MySQL struct {
	Host     string `mapstructure:"host" json:"host"`         // 数据库地址
	Port     int    `mapstructure:"port" json:"port"`         // 数据库端口
	User     string `mapstructure:"user" json:"user"`         // 数据库用户名
	Password string `mapstructure:"password" json:"password"` // 数据库密码
	DBName   string `mapstructure:"dbname" json:"dbname"`     // 数据库名称
}
