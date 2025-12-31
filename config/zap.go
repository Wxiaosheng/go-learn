package config

type Zap struct {
	Level        string `mapstructure:"level"`          // 日志级别
	Prefix       string `mapstructure:"prefix"`         // 日志前缀
	Format       string `mapstructure:"format"`         // 日志格式
	Director     string `mapstructure:"director"`       // 日志存放目录
	EncodeLevel  string `mapstructure:"encode-level"`   // 日志级别编码
	ShowLine     bool   `mapstructure:"show-line"`      // 是否显示行号
	LogInConsole bool   `mapstructure:"log-in-console"` // 是否在控制台输出
}
