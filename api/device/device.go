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
	UserId     int    `p:"user_id" v:"required#用户id不能为空" dc:"user_id"`
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
	Id     int `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"ID"`
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

type SearchByUidReq struct {
	g.Meta `path:"/search_by_uid" method:"get"`
	UserId int `p:"user_id" v:"required|integer|min:1#user_id不能为空|id只能是整数|最小值不应小于1" dc:"用户id"`
}

type SearchByUidRes struct {
	Code          int             `json:"code" dc:"是否有返回值"`
	TencentDevice []entity.Device `json:"tencent_device" dc:"腾讯云设备"`
	AliDevice     []entity.Device `json:"ali_device" dc:"阿里云设备"`
	HuaweiDevice  []entity.Device `json:"huawei_device" dc:"华为云设备"`
}

type SearchByPlatformReq struct {
	g.Meta   `path:"/search_by_platform" method:"get"`
	Platform string `p:"platform" v:"required#平台名不能为空" dc:"平台"`
	UserId   int    `p:"userid" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"用户id不能为空"`
}

type SearchByPlatformRes struct {
	Code    int             `json:"code" dc:"是否有返回值"`
	Devices []entity.Device `json:"devices" dc:"设备列表"`
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

// PubReq 信息上传
type PubReq struct {
	g.Meta   `path:"/publish" method:"post"`
	UserId   int    `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"user_id"`
	TopicId  int    `p:"topic_id" v:"required|integer|min:1#topicid不能为空|id只能是整数|最小值不应小于1" dc:"topic_id"`
	DeviceId int    `p:"device_id" v:"required|integer|integer|min:1#deviceid不能为空|id只能是整数|最小值不应小于1" dc:"device_id"`
	JsonInfo string `p:"json_info" v:"required#json不能为空" dc:"json"`
}

type PubRes struct {
	Code int `json:"code" dc:"返回状态"`
}

// SubReq 订阅topic
type SubReq struct {
	g.Meta   `path:"/subscribe" method:"get"`
	TopicId  int `p:"topic_id" v:"required|integer|min:1#topicid不能为空|id只能是整数|最小值不应小于1" dc:"topic_id"`
	DeviceId int `p:"device_id" v:"required|integer|integer|min:1#deviceid不能为空|id只能是整数|最小值不应小于1" dc:"device_id"`
}

type SubRes struct {
	Code int `json:"code" dc:"返回状态"`
}

// SubInfoReq 获取订阅topic返回信息
type SubInfoReq struct {
	g.Meta   `path:"/subscribe_info" method:"get"`
	TopicId  int `p:"topic_id" v:"required|integer|min:1#topicid不能为空|id只能是整数|最小值不应小于1" dc:"topic_id"`
	DeviceId int `p:"device_id" v:"required|integer|integer|min:1#deviceid不能为空|id只能是整数|最小值不应小于1" dc:"device_id"`
}

type SubInfoRes struct {
	Info string `json:"info" dc:"信息"`
}

type ChatInfoReq struct {
	g.Meta `path:"/chat_info" method:"get"`
	UserId int `p:"user_id" v:"required|integer|min:1#userid不能为空|id只能是整数|最小值不应小于1" dc:"user_id"`
	// 传多少天返回多少天数据
	Days int `p:"days" v:"required|integer|min:1#days不能为空|day只能是整数|最小值不应小于1" dc:"days"`
}

type ChatInfoRes struct {
	Code               int      `json:"code" dc:"是否有返回值"`
	LineChatXData      []string `json:"lineChatXData" dc:"线性表展示"`
	LineChatSeriesData []int    `json:"lineChatSeriesData" dc:"线性表数据"`
	BarChatOnline      []int    `json:"barChatOnline" dc:"柱状图在线数据"`
	BarChatOffline     []int    `json:"barChatOffline" dc:"柱状图离线数据"`
}

type SearchSubTopicReq struct {
	g.Meta   `path:"/get_sub_topic" method:"get"`
	UserId   int `p:"user_id" v:"required|integer|min:1#userid不能为空|id只能是整数|最小值不应小于1" dc:"user_id"`
	DeviceId int `p:"device_id" v:"required|integer|integer|min:1#deviceid不能为空|id只能是整数|最小值不应小于1" dc:"device_id"`
}

type SearchSubTopicRes struct {
	Code     int               `json:"code" dc:"是否有返回值"`
	SubTopic []entity.SubTopic `json:"sub_topic" dc:"订阅的主题"`
}
