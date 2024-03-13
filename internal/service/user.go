package service

import (
	"context"
	"iot_device_simulation/internal/model/do"
	"iot_device_simulation/internal/model/entity"
)

type IUser interface {
	Login(ctx context.Context, username string, password string) (user *entity.User, err error)
	Register(ctx context.Context, user *do.User) (id int, err error)
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("IUser接口未实现或未注册")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
