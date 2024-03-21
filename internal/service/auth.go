package service

import (
	"errors"
	"server/internal/e"
	"server/internal/global"
	"server/internal/model"
	"server/internal/util"
)

type AuthService struct{}

var (
	AuthServiceApp = new(AuthService)
)

func (as *AuthService) Login(u model.User) (userInter *model.User, err error) {
	if global.DB == nil {
		return nil, e.ErrDatabaseNotInitialized
	}

	var user model.User
	if err = global.DB.Where("username = ?", u.Username).First(&user).Error; err != nil {
		return nil, err
	} else {
		if ok := util.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		} else {
			return &user, nil
		}
	}
}
