package config

// ServerConfig
// @Description: 后端配置
type ServerConfig struct {
	// 系统配置
	JwtConfig    JwtConfig    `mapstructure:"jwt"`    // jwt 配置
	SystemConfig SystemConfig `mapstructure:"system"` // 系统配置
	ZapConfig    ZapConfig    `mapstructure:"zap"`    // zap 配置

	// 数据库配置
	MysqlConfig    MysqlConfig    `mapstructure:"mysql"`    // mysql 配置
	PostgresConfig PostgresConfig `mapstructure:"postgres"` // postgres 配置
}
