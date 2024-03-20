package database

import (
	"errors"
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
	if ic.isInitData() {
		return nil
	}
	entities := []gormadapter.CasbinRule{
		{Ptype: "p", V0: "1", V1: "/user/postuserinfo", V2: "POST"},
	}
	if err = global.DB.Create(entities).Error; err != nil {
		return errors.New("权限表初始化失败")
	}
	global.Logger.Info("权限表初始化成功")
	return nil
}

func (ic *initCasbin) isInitData() bool {
	if err := global.DB.Where(gormadapter.CasbinRule{Ptype: "p", V0: "1", V1: "/api/v1/user/getuserinfo", V2: "POST"}).First(&gormadapter.CasbinRule{}).Error; err != nil {
		return false
	} else {
		return true
	}
}
