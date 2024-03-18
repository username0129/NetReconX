package config

// PostgresConfig
// @Description: PostgreSQL 配置文件
type PostgresConfig struct {
	DatabaseGeneralConfig `mapstructure:",squash"` // 结构体嵌入
}
