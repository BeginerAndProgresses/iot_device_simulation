package user

import (
	"context"
	"errors"
	"iot_device_simulation/internal/dao"
	"iot_device_simulation/internal/model/do"
	"iot_device_simulation/internal/model/entity"
	"iot_device_simulation/internal/service"
)

func init() {
	service.RegisterUser(New())
}

type iUser struct {
}

func New() *iUser {
	return &iUser{}
}

func (i *iUser) Login(ctx context.Context, username string, password string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(do.User{
		Username: username,
		Password: password,
	}).Scan(&user)

	if user == nil {
		err = errors.New("用户名或密码有误")
	}
	return
}

func (i *iUser) Register(ctx context.Context, user *do.User) (id int, err error) {
	result, err := dao.User.Ctx(ctx).Data(user).Insert()
	insertId, err := result.LastInsertId()
	id = int(insertId)
	if err != nil {
		err = errors.New("注册失败")
	}
	return
}

func (i *iUser) Search(ctx context.Context, id int) (user entity.User, err error) {
	err = dao.User.Ctx(ctx).Where("id", id).Scan(&user)
	return
}

func (i *iUser) Update(ctx context.Context, newUser entity.User) (err error) {
	_, err = dao.User.Ctx(ctx).Where("id", newUser.Id).Data(&newUser).OmitEmptyData().Update()
	return
}
