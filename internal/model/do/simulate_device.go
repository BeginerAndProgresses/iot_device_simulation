// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SimulateDevice is the golang structure of table simulate_device for DAO operations like Where/Data.
type SimulateDevice struct {
	g.Meta     `orm:"table:simulate_device, do:true"`
	Id         interface{} //
	PlatForm   interface{} // 平台
	DeviceName interface{} // 设备名
	DeviceId   interface{} // 设备id
	State      interface{} // 状态
	ProductId  interface{} // 产品id
	UserId     interface{} // 用户id
	Intervals  interface{} // 时间间隔，默认20秒
	MqttOption interface{} // mqtt连接参数
	Topic      interface{} // 通信主题
}
