package config

// ZapConfig
// @Description: zap 日志配置
type ZapConfig struct {
	Level        string `mapstructure:"level"`          // 最低记录日志等级
	Format       string `mapstructure:"format"`         // 输出格式
	Director     string `mapstructure:"director"`       // 日志保存路径
	EncodeLevel  string `mapstructure:"encode_level"`   // 日志级别的编码器
	LogInConsole bool   `mapstructure:"log_in_console"` // 日志输出到控制台
}
