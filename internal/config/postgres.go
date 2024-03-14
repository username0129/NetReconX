package config

type PostgresConfig struct {
	DatabaseGeneralConfig `mapstructure:",squash"` // 结构体嵌入
}
