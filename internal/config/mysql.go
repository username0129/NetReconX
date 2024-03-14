package config

import "fmt"

type MysqlConfig struct {
	DatabaseGeneralConfig `mapstructure:",squash"` // 结构体嵌入
}

// GetDSN
//
//	@Description: 获取用于连接 MySQL 数据库的 DSN
//	@receiver mysqlConfig
//	@return string
func (mysqlConfig *MysqlConfig) GetDSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DBName)
}

// GetEmptyDSN
//
//	@Description: 获取用于连接 MySQL 的 DSN
//	@receiver mysqlConfig
//	@return string
func (mysqlConfig *MysqlConfig) GetEmptyDSN() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port)
}
