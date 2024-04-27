package device

import (
	"github.com/gogf/gf/v2/frame/g"
	"iot_device_simulation/internal/model/entity"
)

type AddSimulateDeviceReq struct {
	g.Meta         `path:"/add_simulate_device" method:"post" tags:"模拟设备"`
	UserId         int    `p:"user_id" v:"required|integer|min:1#userid不能为空|id只能是整数|最小值不应小于1" dc:"用户id"`
	Platform       string `p:"platform" v:"required#平台名不能为空" dc:"平台"`
	Topic          string `p:"topic" v:"required#topic不能为空" dc:"topic"`
	Interval       int64  `p:"interval" v:"required|integer|min:1#interval不能为空|id只能是整数|最小值不应小于1" dc:"时间间隔"`
	DeviceName     string `p:"device_name" v:"required#设备名不能为空" dc:"设备名"`
	DeviceId       string `p:"device_id" v:"required#设备id不能为空" dc:"设备id"`
	ProductId      string `p:"product_id" v:"required#产品ID不能为空" dc:"产品id"`
	ClientId       string `p:"client_id" v:"required#client_id不能为空" dc:"client_id"`
	Port           int    `p:"port" v:"required|integer|min:1#port不能为空|port只能是整数|最小值不应小于1" dc:"port"`
	ServiceAddress string `p:"service_address" v:"required#service_address不能为空" dc:"service_address"`
	UserName       string `p:"user_name" v:"required#user_name不能为空" dc:"user_name"`
	Password       string `p:"password" v:"required#password不能为空" dc:"password"`
}

type AddSimulateDeviceRes struct {
	Code int `json:"code" dc:"返回状态"`
}

type ConnSimulateDeviceReq struct {
	g.Meta   `path:"/conn_simulate_device" method:"get" tags:"模拟设备"`
	DeviceId int `p:"device_id" v:"required|integer|integer|min:1#deviceid不能为空|id只能是整数|最小值不应小于1" dc:"device_id"`
}

type ConnSimulateDeviceRes struct {
	State int `json:"state" dc:"状态"`
}

type DisConnSimulateDeviceReq struct {
	g.Meta   `path:"/dis_conn_simulate_device" method:"get" tags:"模拟设备"`
	DeviceId int `p:"device_id" v:"required|integer|integer|min:1#deviceid不能为空|id只能是整数|最小值不应小于1" dc:"device_id"`
}

type DisConnSimulateDeviceRes struct {
	State int `json:"state" dc:"状态"`
}

type GetSimulateDeviceByUidReq struct {
	g.Meta `path:"/get_simulate_device_by_uid" method:"get" tags:"模拟设备"`
	UserId int `p:"user_id" v:"required|integer|min:1#userid不能为空|id只能是整数|最小值不应小于1" dc:"用户id"`
}

type GetSimulateDeviceByUidRes struct {
	Code           int                     `json:"code" dc:"是否有返回值"`
	SimulateDevice []entity.SimulateDevice `json:"simulate_device" dc:"模拟设备"`
}

type EditSimulateDeviceReq struct {
	g.Meta         `path:"/edit_simulate_device" method:"put" tags:"模拟设备"`
	Id             int    `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"id"`
	Interval       int64  `p:"interval" v:"required|integer|min:1#interval不能为空|id只能是整数|最小值不应小于1" dc:"时间间隔"`
	DeviceId       string `p:"device_id" v:"required#设备id不能为空" dc:"设备id"`
	ClientId       string `p:"client_id" v:"required#client_id不能为空" dc:"client_id"`
	Port           int    `p:"port" v:"required|integer|min:1#port不能为空|port只能是整数|最小值不应小于1" dc:"port"`
	ServiceAddress string `p:"service_address" v:"required#service_address不能为空" dc:"service_address"`
	UserName       string `p:"user_name" v:"required#user_name不能为空" dc:"user_name"`
	Password       string `p:"password" v:"required#password不能为空" dc:"password"`
}

type EditSimulateDeviceRes struct {
	Code int `json:"code" dc:"返回状态"`
}

type DeleteSimulateDeviceReq struct {
	g.Meta `path:"/delete_simulate_device" method:"delete" tags:"模拟设备"`
	Id     int `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"id"`
}

type DeleteSimulateDeviceRes struct {
	Code int `json:"code" dc:"返回状态"`
}
