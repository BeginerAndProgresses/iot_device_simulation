package user

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"iot_device_simulation/api/user"
	"iot_device_simulation/internal/consts"
	"iot_device_simulation/internal/model/do"
	"iot_device_simulation/internal/model/entity"
	"iot_device_simulation/internal/service"
	"time"
)

var UserController = &cUser{}

type cUser struct {
}

func (c *cUser) Login(ctx context.Context, req *user.LoginReq) (res *user.LoginRes, err error) {
	en_user, err := service.User().Login(ctx, req.Username, req.Password)
	if err == nil {
		res = &user.LoginRes{
			Token: jwtToken(en_user),
			UserInfo: &user.UserInfo{
				Id:       en_user.Id,
				Username: en_user.Username,
				Nickname: en_user.Nikename,
				Avatar:   en_user.Avatar,
			},
		}
	}
	return
}

// 过期时间为一天
func jwtToken(user *entity.User) string {
	// 实际使用中可将Key存于文件中或放在常量中 consts.JwtTokenKey

	claim := jwt.RegisteredClaims{
		Subject:   user.Username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}
	//claim := jwt.MapClaims{
	//	"Data":      entity.Device{Id: user.Id},
	//	"ExpiresAt": jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	//}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(consts.JwtTokenKey))

	if err != nil {
		panic("token生成出错")
	}

	return token
}

func (c *cUser) Register(ctx context.Context, req *user.RegisterReq) (res *user.RegisterRes, err error) {
	id, err := service.User().Register(ctx, &do.User{Username: req.Username, Password: req.Password})
	if err == nil {
		res = &user.RegisterRes{
			Id: id,
		}
	}
	return
}

func (c *cUser) SearchUser(ctx context.Context, req *user.SearchReq) (res *user.SearchRes, err error) {
	search, err := service.User().Search(ctx, req.Id)
	res = &user.SearchRes{
		User: search,
	}
	if err != nil {
		res.Code = 0
	}
	res.Code = 1
	return
}

func (c *cUser) Update(ctx context.Context, req *user.UpdateReq) (res *user.UpdateRes, err error) {
	err = service.User().Update(ctx, entity.User{Id: req.Id, Avatar: req.AvatarUrl, Nikename: req.NikeName})
	return
}
