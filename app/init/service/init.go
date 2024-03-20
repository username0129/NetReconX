package service

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"server/app/init/model"
	"server/internal/e"
	"server/internal/global"
	"sort"
)

// -------------------- 数据库初始化 ----------------------------

// DBInitializer 数据库初始化流程
type DBInitializer interface {
	CreateDatabase(c context.Context, req model.InitRequest) (context.Context, error) // 初始化数据库
	CreateTable(o initSlice) error                                                    // 初始化数据表
	InitData(o initSlice) error
	WriteConfig(c context.Context) error // 写入到配置文件
}

type InitService struct{}

func (is *InitService) Init(req model.InitRequest) (err error) {

	if len(dataInitializers) == 0 {
		return e.ErrNoInitializersAvailable
	}
	sort.Sort(&dataInitializers) // 对初始化进行排序，确保执行顺序

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

	if c, err = dbInitializer.CreateDatabase(c, req); err != nil {
		return err
	} // 创建数据库

	db := c.Value("db").(*gorm.DB)
	global.DB = db

	if err = dbInitializer.CreateTable(dataInitializers); err != nil {
		return err
	}
	if err = dbInitializer.InitData(dataInitializers); err != nil {
		return err
	}
	if err = dbInitializer.WriteConfig(c); err != nil {
		return err
	}
	return nil
}

func createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

// -------------------- 数据表初始化 ----------------------------

// DataInitializer 数据初始化
type DataInitializer interface {
	CreateTable() (err error) // 创建数据表
	InitData() (err error)    // 初始化数据
}

// OrderedDataInitializer 数据初始化流程，通过排序确定初始化顺序
type OrderedDataInitializer struct {
	Order int
	DataInitializer
}

// 实现  sort interface 接口
type initSlice []*OrderedDataInitializer

func (is initSlice) Len() int {
	return len(is)
}

func (is initSlice) Less(i, j int) bool {
	return is[i].Order < is[j].Order
}

func (is initSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

var (
	dataInitializers initSlice
)

func RegisterDataInitializers(order int, initializer DataInitializer) {
	if dataInitializers == nil {
		dataInitializers = []*OrderedDataInitializer{}
	}
	dataInitializer := &OrderedDataInitializer{
		Order:           order,
		DataInitializer: initializer,
	}
	dataInitializers = append(dataInitializers, dataInitializer)
}

func createTable(o []*OrderedDataInitializer) error {
	for _, initializer := range o {
		if err := initializer.CreateTable(); err != nil {
			return err
		}
	}
	return nil
}
