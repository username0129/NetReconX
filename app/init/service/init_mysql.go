package service

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server/app/init/model"
	"server/internal/config"
)

type MysqlInitializer struct{}

func NewMysqlInitializer() *MysqlInitializer {
	return &MysqlInitializer{}
}

// InitDatabase
//
//	@Description: 初始化数据库
//	@receiver mi
//	@param c
//	@param req
//	@return context.Context
//	@return error
func (mi *MysqlInitializer) InitDatabase(c context.Context, req model.InitRequest) (context.Context, error) {
	mysqlConfig := req.ToMySQLConfig()
	c = context.WithValue(c, "config", mysqlConfig)
	if mysqlConfig.DBName == "" {
		return c, nil
	} // 库名为空，不进行数据初始化

	dsn := mysqlConfig.GetEmptyDSN()
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", mysqlConfig.DBName)
	if err := CreateDatabase(dsn, "mysql", createSql); err != nil {
		return nil, err
	} // 创建数据库

	if db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysqlConfig.GetDSN(),
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用自动创建外键约束
	}); err != nil {
		return c, err
	} else {
		c = context.WithValue(c, "db", db)
		return c, nil
	}
}

func (mi *MysqlInitializer) WriteConfig(c context.Context) (err error) {
	mysqlConfig := c.Value("config").(config.MysqlConfig)
	config.GlobalConfig.SystemConfig.DBType = "mysql"
	config.GlobalConfig.MysqlConfig = mysqlConfig

	config.GlobalViper.Set("system.db_type", "mysql")

	// 写入修改后的配置到新文件
	if err = viper.WriteConfigAs(config.GlobalConfig.SystemConfig.ConfigPath); err != nil {
		return err
	}
	return nil
}
