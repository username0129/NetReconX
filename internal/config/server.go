package config

// ServerConfig
// @Description: 后端配置
type ServerConfig struct {
	// 系统配置
	JwtConfig    JwtConfig    `mapstructure:"jwt" yaml:"jwt" json:"jwt"`          // jwt 配置
	SystemConfig SystemConfig `mapstructure:"system" yaml:"system" json:"system"` // 系统配置
	ZapConfig    ZapConfig    `mapstructure:"zap" yaml:"zap" json:"zap"`          // zap 配置

	// 数据库配置
	MysqlConfig MysqlConfig `mapstructure:"mysql" yaml:"mysql" json:"mysql"` // mysql 配置
}
