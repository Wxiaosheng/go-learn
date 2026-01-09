package config

import (
	"fmt"
)

type MySQL struct {
	Host         string `mapstructure:"host" json:"host"`         // 数据库地址
	Port         int    `mapstructure:"port" json:"port"`         // 数据库端口
	User         string `mapstructure:"user" json:"user"`         // 数据库用户名
	Password     string `mapstructure:"password" json:"password"` // 数据库密码
	DBName       string `mapstructure:"dbname" json:"dbname"`     // 数据库名称
	Charset      string `mapstructure:"charset" json:"charset"`   // 字符集
	Timeout      int    `mapstructure:"timeout" json:"timeout"`   // 连接超时时间，单位秒
	MaxIdleConns int    `mapstructure:"max-idle" json:"max-idle"` // 最大空闲连接数
	MaxOpenConns int    `mapstructure:"max-open" json:"max-open"` // 最大连接数
}

/* 获取 DSN 连接字符串 */
func (m *MySQL) DSN() string {
	charset := m.Charset
	if charset == "" {
		charset = "utf8mb4" // 默认字符集
	}
	// Timeout is specified in seconds in the config
	timeout := m.Timeout
	if timeout == 0 {
		timeout = 30 // 默认连接超时 30s
	}

	fmt.Printf(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&timeout=%ds",
		m.User, m.Password, m.Host, m.Port, m.DBName, charset, timeout))

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&timeout=%ds",
		m.User, m.Password, m.Host, m.Port, m.DBName, charset, timeout)
}
