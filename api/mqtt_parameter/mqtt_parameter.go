package mqtt_parameter

import (
	"github.com/gogf/gf/v2/frame/g"
	"iot_device_simulation/internal/model/entity"
)

type AddReq struct {
	g.Meta        `path:"/" method:"post"`
	DeviceId      int    `p:"device_id" v:"required#id不能为空" dc:"id"`
	ClientId      string `p:"client_id" v:"required#client_id不能为空" dc:"client_id"`
	Port          int    `p:"port" v:"required|integer#端口不能为空|post类型应为整数" dc:"端口"`
	ServerAddress string `p:"server_address" v:"required#服务器地址不能为空" dc:"服务器地址"`
	Username      string `p:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password      string `p:"password" v:"required#密码不能为空" dc:"密码"`
	UserId        int    `p:"user_id" v:"required|integer#user_id不能为空|user_id应为整数" dc:"user_id"`
}

type AddRes struct {
	Id int `json:"id" dc:"连接参数id"`
}

type EditReq struct {
	g.Meta        `path:"/" method:"put"`
	Id            int    `p:"id" v:"required|integer#id不能为空|id应为整数" dc:"id"`
	ClientId      string `p:"client_id"  dc:"client_id"`
	Port          int    `p:"port" v:"integer#post应为整数" dc:"端口"`
	ServerAddress string `p:"server_address"  dc:"服务器地址"`
	Username      string `p:"username" dc:"用户名"`
	Password      string `p:"password" dc:"密码"`
	UserId        int    `p:"user_id" v:"required|integer#user_id不能为空|user_id应为整数" dc:"user_id"`
}

type EditRes struct {
	Id int `json:"id" dc:"连接参数id"`
}

type DelReq struct {
	g.Meta `path:"/" method:"delete"`
	Id     int `p:"id" v:"required|integer#id不能为空|id应为整数" dc:"id"`
}

type DelRes struct {
	Id int `json:"id" dc:"连接参数id"`
}

type SearchReq struct {
	g.Meta `path:"/" method:"get"`
	Id     int `p:"id" v:"required|integer#id不能为空|id应为整数" dc:"id"`
}

type SearchRes struct {
	Code int                  `json:"code" dc:"是否有值"`
	Mqtt entity.MqttParameter `json:"mqtt" dc:"实体"`
}
