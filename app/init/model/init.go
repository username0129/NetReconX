package model

import (
	"server/internal/config"
	"time"
)

type InitRequest struct {
	DBType   string `json:"db_type,omitempty" yaml:"db_type"`   // 数据库类型
	Host     string `json:"host,omitempty" yaml:"host"`         // 服务器地址
	Port     string `json:"port,omitempty" yaml:"port"`         // 数据库连接端口
	DBName   string `json:"db_name,omitempty" yaml:"db_name"`   // 数据库名
	Username string `json:"username,omitempty" yaml:"username"` // 数据库用户名
	Password string `json:"password,omitempty" yaml:"password"` // 数据库密码
}

// ToMySQLConfig
//
//	@Description: 转换为 Mysql 配置
//	@receiver ir
//	@return config.MysqlConfig
func (ir *InitRequest) ToMySQLConfig() config.MysqlConfig {
	return config.MysqlConfig{
		DatabaseGeneralConfig: config.DatabaseGeneralConfig{
			Host:            ir.Host,
			Port:            ir.Port,
			DBName:          ir.DBName,
			Username:        ir.Username,
			Password:        ir.Password,
			MaxIdleConns:    10,
			MaxOpenConns:    100,
			ConnMaxLifetime: 10 * time.Second,
		},
	}
}
