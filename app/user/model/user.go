package model

import (
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID        uuid.UUID `mapstructure:"uuid" gorm:"index;comment:用户 UUID"`
	Username    string    `mapstructure:"username" gorm:"index;comment:用户登录名"`
	Password    string    `mapstructure:"password" gorm:"comment:用户登录密码"`
	AuthorityId uint      `mapstructure:"authority_id" gorm:"default:0;comment:用户身份 ID"`
	Phone       string    `mapstructure:"phone" gorm:"comment:手机号"`
	Email       string    `mapstructure:"email" gorm:"comment:邮箱"`
	Enable      int       `mapstructure:"enable" gorm:"default:1;comment:用户状态 1正常 0冻结"`
}
