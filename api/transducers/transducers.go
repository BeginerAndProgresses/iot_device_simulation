package transducers

import (
	"github.com/gogf/gf/v2/frame/g"
	"iot_device_simulation/internal/model/entity"
)

type AddTransducersReq struct {
	g.Meta          `path:"/" method:"post"  tag:"传感器"`
	UserId          int    `p:"user_id" v:"required|integer|min:1#user_id不能为空|id只能是整数|最小值不应小于1" dc:"user_id"`
	Identifier      string `p:"identifier" v:"required|length:1,32#identifier不能为空|identifier长度为1~32位" dc:"标识符"`
	TransducersType string `p:"transducers_type" v:"required|length:1,32#transducers_type不能为空|transducers_type长度为1~32位" dc:"传感器类型"`
	Options         g.Map  `p:"options" v:"required|json#options不能为空|options格式不正确" dc:"options"`
}

type AddTransducersRes struct {
	Code int `json:"code" dc:"返回状态"`
}

type GetTransducersByDeviceIdReq struct {
	g.Meta   `path:"/by_device_id" method:"get" tag:"传感器"`
	DeviceId int `p:"device_id" v:"required|integer|min:1#device_id不能为空|id只能是整数|最小值不应小于1" dc:"device_id"`
}

type GetTransducersByDeviceIdRes struct {
	Code        int                  `json:"code" dc:"返回状态"`
	Transducers []entity.Transducers `json:"transducers" dc:"transducers"`
}

type GetTransducersByUserIdReq struct {
	g.Meta `path:"/by_user_id" method:"get" tag:"传感器"`
	UserId int `p:"user_id" v:"required|integer|min:1#user_id不能为空|id只能是整数|最小值不应小于1" dc:"user_id"`
	Page   int `p:"page" v:"required|integer|min:1#page不能为空|page只能是整数|最小值不应小于1" dc:"page"`
	Size   int `p:"size" v:"required|integer|min:1#size不能为空|size只能是整数|最小值不应小于1" dc:"size"`
}

type GetTransducersByUserIdRes struct {
	Code        int                  `json:"code" dc:"返回状态"`
	Transducers []entity.Transducers `json:"transducers" dc:"transducers"`
	Count       int                  `json:"count" dc:"count"`
}

type GetTransducersByUidButNotDeviceReq struct {
	g.Meta `path:"/by_uid_but_not_device" method:"get" tag:"传感器"`
	UserId int `p:"user_id" v:"required|integer|min:1#user_id不能为空|id只能是整数|最小值不应小于1" dc:"user_id"`
}

type GetTransducersByUidButNotDeviceRes struct {
	Code        int                  `json:"code" dc:"返回状态"`
	Transducers []entity.Transducers `json:"transducers" dc:"transducers"`
}

type GetTransducerByIdReq struct {
	g.Meta `path:"/" method:"get" tag:"传感器"`
	Id     int `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"id"`
}

type GetTransducerByIdRes struct {
	Code        int                `json:"code" dc:"返回状态"`
	Transducers entity.Transducers `json:"transducers" dc:"transducers"`
}

type UpdateTransducersByDeviceIdReq struct {
	g.Meta     `path:"/" method:"put" tag:"传感器"`
	Id         int    `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"id"`
	Identifier string `p:"identifier" v:"required|length:1,32#identifier不能为空|identifier长度为1~32位" dc:"标识符"`
	Options    g.Map  `p:"options" v:"required|json#options不能为空|options格式不正确" dc:"options"`
}

type UpdateTransducersByDeviceIdRes struct {
	Code int `json:"code" dc:"返回状态"`
}

type DeleteTransducersByIdReq struct {
	g.Meta `path:"/" method:"delete" tag:"传感器"`
	Id     int `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"id"`
}

type DeleteTransducersByIdRes struct {
	Code int `json:"code" dc:"返回状态"`
}

type SetDeviceIdReq struct {
	g.Meta   `path:"/set_deviceId" method:"put" tag:"传感器"`
	Id       int `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"id"`
	DeviceId int `p:"device_id" v:"required|integer|min:1#device_id不能为空|id只能是整数|最小值不应小于1" dc:"device_id"`
}

type SetDeviceIdRes struct {
	Code int `json:"code" dc:"返回状态"`
}
