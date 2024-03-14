package database

import (
	"gorm.io/gorm"
	"server/internal/config"
)

// GetDatabaseConnection
//
//	@Description: 使用 gorm 获取数据库链接
//	@return *gorm.DB
func GetDatabaseConnection() *gorm.DB {
	switch config.GlobalConfig.SystemConfig.DBType {
	case "mysql":
		return MysqlConnection()
	case "postgres":
		return PostgresConnection()
	default:
		return MysqlConnection()
	} // 根据系统使用的数据库类型返回对应的数据库链接。默认使用 MySQL
}
