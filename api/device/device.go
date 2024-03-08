package device

import (
	"github.com/gogf/gf/v2/frame/g"
	"iot_device_simulation/internal/model/entity"
)

type AddReq struct {
	g.Meta     `path:"/" method:"post"`
	Platform   string `p:"platform" v:"required#平台不能为空" dc:"平台"`
	DeviceName string `p:"device_name" v:"required#设备名不能为空" dc:"设备名"`
	ProductId  string `p:"product_id" v:"required#产品ID不能为空" dc:"产品id"`
}

type AddRes struct {
	Id int `json:"id" dc:"设备id"`
}

type SearchReq struct {
	g.Meta `path:"/" method:"get"`
	Id     int `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"平台"`
}

type SearchRes struct {
	Code   int           `json:"code" dc:"是否有返回值"`
	Device entity.Device `json:"device" dc:"设备"`
}

type DelReq struct {
	g.Meta `path:"/" method:"delete"`
	Id     int `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"平台"`
}

type DelRes struct {
	Id int `json:"id" dc:"设备id"`
}

type EditReq struct {
	g.Meta          `path:"/" method:"put"`
	Id              int    `p:"id" v:"integer|min:1#id只能是整数|最小值不应小于1" dc:"id"`
	Platform        string `p:"platform" dc:"平台"`
	DeviceName      string `p:"device_name" dc:"设备名"`
	MqttParameterId uint   `p:"mqtt_parameter_id" dc:"连接id"`
	ProductId       string `p:"product_id" dc:"产品id"`
}

type EditRes struct {
	Id int `json:"id" dc:"设备id"`
}

type ConnReq struct {
	g.Meta `path:"/conn" method:"put"`
	Id     int `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"id"`
}

type ConnRes struct {
	Id    int `json:"id" dc:"id"`
	State int `json:"state" dc:"状态"` // 0 未连接，1 连接中 ，2 未设置链接参数
}

type DisConnReq struct {
	g.Meta `path:"/dis_conn" method:"put"`
	Id     int `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"id"`
}

type DisConnRes struct {
	Id    int `json:"id" dc:"id"`
	State int `json:"state" dc:"状态"` // 1 未连接，2 连接中 ，3 未设置链接参数
}

// 属性上传
type PubReq struct {
	g.Meta   `path:"/publish" method:"post"`
	TopicId  int    `p:"topic_id" v:"required|integer|min:1#topicid不能为空|id只能是整数|最小值不应小于1" dc:"topic_id"`
	DeviceId int    `p:"device_id" v:"required|integer|integer|min:1#deviceid不能为空|id只能是整数|最小值不应小于1" dc:"device_id"`
	JsonInfo string `p:"json_info" v:"required#json不能为空" dc:"json"`
}

type PubRes struct {
	Code int `json:"code" dc:"返回状态"`
}
