package user

import "github.com/gogf/gf/v2/frame/g"

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
