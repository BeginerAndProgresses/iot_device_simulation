package user

import (
	"github.com/gogf/gf/v2/frame/g"
	"iot_device_simulation/internal/model/entity"
)

type LoginReq struct {
	g.Meta   `path:"/login" method:"post"`
	Username string `p:"username" v:"required#请输入用户名" dc:"用户名"`
	Password string `p:"password" v:"required#请输入密码" dc:"密码"`
}

type UserInfo struct {
	Id       int    `json:"id"         dc:"用户ID"`
	Username string `json:"username"   dc:"用户名"`
	Nickname string `json:"nickname"   dc:"昵称"`
	Avatar   string `json:"avatar"     dc:"用户头像"`
}

type LoginRes struct {
	Token    string    `json:"token" dc:"验证token"`
	UserInfo *UserInfo `json:"user_info"`
}

type RegisterReq struct {
	g.Meta   `path:"/register" method:"post"`
	Username string `p:"username" v:"required#请输入用户名" dc:"用户名"`
	Password string `p:"password" v:"required#请输入密码" dc:"密码"`
}

type RegisterRes struct {
	Id int `json:"id" dc:"返回0失败，否则成功"`
}

type SearchReq struct {
	g.Meta `path:"/search" method:"get"`
	Id     int `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"ID"`
}

type SearchRes struct {
	Code int         `json:"code" dc:"状态码"`
	User entity.User `json:"user" dc:"用户"`
}

type UpdateReq struct {
	g.Meta    `path:"/update" method:"put"`
	Id        int    `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"ID"`
	NikeName  string `p:"nike_name" v:"required#昵称不能为空" dc:"昵称"`
	AvatarUrl string `p:"avatar_url" v:"required#url不能为空" dc:"头像"`
}

type UpdateRes struct {
}
