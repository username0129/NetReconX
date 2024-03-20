package model

import (
	"server/internal/config"
	"server/internal/global"
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
//	@return config.Mysql
func (ir *InitRequest) ToMySQLConfig() config.Mysql {
	cfg := global.Config.Mysql
	cfg.Host = ir.Host
	cfg.Port = ir.Port
	cfg.DBName = ir.DBName
	cfg.Username = ir.Username
	cfg.Password = ir.Password
	return cfg
}
