package config

import (
	"fmt"
	"time"
)

// MysqlConfig
// @Description: MySQL 配置文件
type MysqlConfig struct {
	Host            string        `mapstructure:"host" yaml:"host" json:"host,omitempty"`                                        // 数据库主机地址
	Port            string        `mapstructure:"port" yaml:"port" json:"port,omitempty"`                                        // 数据库所用端口
	DBName          string        `mapstructure:"db_name" yaml:"db_name" json:"db_name,omitempty"`                               // 数据库名
	Username        string        `mapstructure:"username" yaml:"username" json:"username,omitempty"`                            // 数据库用户名
	Password        string        `mapstructure:"password" yaml:"password" json:"password,omitempty"`                            // 数据库密码
	MaxIdleConns    int           `mapstructure:"max_idle_conns" yaml:"max_idle_conns" json:"max_idle_conns,omitempty"`          // 空闲连接最大数量
	MaxOpenConns    int           `mapstructure:"max_open_conns" yaml:"max_open_conns" json:"max_open_conns,omitempty"`          // 打开连接的最大数量
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime" yaml:"conn_max_lifetime" json:"conn_max_lifetime,omitempty"` // 空闲连接存活时间
}

// GetDSN
//
//	@Description: 获取用于连接 MySQL 数据库的 DSN
//	@receiver mysqlConfig
//	@return string
func (mc *MysqlConfig) GetDSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", mc.Username, mc.Password, mc.Host, mc.Port, mc.DBName)
}

// GetEmptyDSN
//
//	@Description: 获取用于连接 MySQL 的 DSN
//	@receiver mysqlConfig
//	@return string
func (mc *MysqlConfig) GetEmptyDSN() string {
	if mc.Host == "" {
		mc.Host = "127.0.0.1"
	}
	if mc.Port == "" {
		mc.Port = "3306"
	}
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/", mc.Username, mc.Password, mc.Host, mc.Port)
}
