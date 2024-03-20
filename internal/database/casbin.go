package database

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"server/app/init/service"
	"server/internal/global"
)

type initCasbin struct{}

func init() {
	service.RegisterDataInitializers(10, &initCasbin{})
}

func (ic *initCasbin) CreateTable() (err error) {
	if !global.DB.Migrator().HasTable(&gormadapter.CasbinRule{}) {
		if err = global.DB.AutoMigrate(&gormadapter.CasbinRule{}); err != nil {
			return err
		}
	}
	return nil
}

func (ic *initCasbin) InitData() (err error) {
	return nil
}
