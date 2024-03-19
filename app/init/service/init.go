package service

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"server/app/init/model"
	"server/internal/config"
)

type DBInitializer interface {
	InitDatabase(c context.Context, req model.InitRequest) (context.Context, error) // 初始化数据库
	WriteConfig(c context.Context) error                                            // 写入到配置文件
}

type InitService struct{}

func (is *InitService) Init(req model.InitRequest) (err error) {
	c := context.TODO()
	var dbInitializer DBInitializer
	switch req.DBType {
	case "mysql":
		dbInitializer = NewMysqlInitializer()
		c = context.WithValue(c, "db_type", "mysql")
	default:
		dbInitializer = NewMysqlInitializer()
		c = context.WithValue(c, "db_type", "mysql")
	} // 获取对应的数据库初始化工具

	if c, err = dbInitializer.InitDatabase(c, req); err != nil {
		return err
	} // 创建数据库

	db := c.Value("db").(*gorm.DB)
	config.GlobalDB = db

	if err = dbInitializer.WriteConfig(c); err != nil {
		return err
	}
	return nil
}

func CreateDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}
