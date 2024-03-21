package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
}

func (Menu) TableName() string {
	return "sys_menus"
}
