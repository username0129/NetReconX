package database

import (
	"errors"
	"github.com/gofrs/uuid/v5"
	"server/app/init/service"
	"server/app/user/model"
	"server/internal/global"
	"server/internal/utils"
)

type initUser struct {
}

func init() {
	service.RegisterDataInitializers(11, &initUser{})
}

func (iu *initUser) CreateTable() (err error) {
	if !global.DB.Migrator().HasTable(&model.User{}) {
		if err = global.DB.AutoMigrate(&model.User{}); err != nil {
			return err
		}
	}
	return nil
}

func (iu *initUser) InitData() (err error) {
	if iu.isInitData() {
		return nil
	}
	adminPassword := utils.BcryptHash("123456")
	entities := []model.User{
		{
			UUID:        uuid.Must(uuid.NewV4()),
			Username:    "admin",
			Password:    adminPassword,
			AuthorityId: uint(1),
		},
	}
	if global.DB.Create(entities).Error != nil {
		return errors.New("用户表初始化失败")
	}
	global.Logger.Info("用户表初始化成功")
	return nil
}

func (iu *initUser) isInitData() bool {
	if err := global.DB.Where("username = ?", "admin").First(&model.User{}).Error; err != nil {
		return false
	} else {
		return true
	}
}
