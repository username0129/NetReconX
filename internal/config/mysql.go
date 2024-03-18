package config

import "fmt"

// MysqlConfig
// @Description: MySQL 配置文件
type MysqlConfig struct {
	DatabaseGeneralConfig `mapstructure:",squash"` // 结构体嵌入
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
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/", mc.Username, mc.Password, mc.Host, mc.Port)
}
