package dbutil

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"server/internal/global"
	"server/internal/model"
	"server/internal/util"
)

// InitialData 用于数据库初始化的数据结构
type InitialData struct {
	TableName string
	Data      []interface{}
}

// initialDatas 定义初始化数据库时使用的数据
var initialDatas = []InitialData{
	{
		TableName: "sys_authorities",
		Data: []interface{}{
			&model.Authority{AuthorityName: "系统管理员"},
		},
	},
	{
		TableName: "sys_users",
		Data: []interface{}{
			&model.User{Username: "admin", Password: util.BcryptHash("123456"), Nickname: "系统管理员", AuthorityId: 1, Authorities: []model.Authority{{AuthorityName: "系统管理员"}}},
		},
	},
}

// CommonDBOperations 定义了数据库操作的公共接口
type CommonDBOperations struct{}

// CreateTable 创建表结构
func (c *CommonDBOperations) CreateTable() error {
	for _, initData := range initialDatas {
		tableName := initData.TableName
		exists := global.DB.Migrator().HasTable(tableName) // 检查表是否存在
		if !exists {
			if err := global.DB.AutoMigrate(initData.Data...); err != nil {
				return fmt.Errorf("创建表 %s 失败: %w", tableName, err)
			}
		}
	}
	return nil
}

func (c *CommonDBOperations) InsertData() error {
	tx := global.DB.Begin() // 回滚事务，避免出现只完成了部分插入的情况。

	for _, initData := range initialDatas {
		for _, data := range initData.Data {
			if initializableData, ok := data.(model.Initializable); ok {
				if err := initializableData.InsertData(global.DB); err != nil {
					tx.Rollback() // 插入失败，回滚事务
					return fmt.Errorf("初始化表 %s 失败: %w", initData.TableName, err)
				}
			} else {
				tx.Rollback() // 类型断言失败，回滚事务
				return fmt.Errorf("数据项 %v 不支持初始化接口", data)
			}

		}
	}
	// 提交事务
	return tx.Commit().Error
}

func ExecuteSQL(dsn string, driver string, sqlStatement string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(sqlStatement)
	return err
}
