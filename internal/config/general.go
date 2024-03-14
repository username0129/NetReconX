package config

import "time"

// DatabaseGeneralConfig
// @Description: 数据库通用配置，方便后续扩展
type DatabaseGeneralConfig struct {
	Host            string        `mapstructure:"host"`              // 数据库主机地址
	Port            string        `mapstructure:"port"`              // 数据库所用端口
	DBName          string        `mapstructure:"db_name"`           // 数据库名
	Username        string        `mapstructure:"username"`          // 数据库用户名
	Password        string        `mapstructure:"password"`          // 数据库密码
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`    // 空闲连接最大数量
	MaxOpenConns    int           `mapstructure:"max_open_conns"`    // 打开连接的最大数量
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"` // 空闲连接存活时间
}
