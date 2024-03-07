// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Device is the golang structure of table device for DAO operations like Where/Data.
type Device struct {
	g.Meta          `orm:"table:device, do:true"`
	Id              interface{} //
	PlatForm        interface{} // 平台名称
	DeviceName      interface{} // 设备名称
	MqttParameterId interface{} //
	State           interface{} // 状态，如果未启动为0，如果启动为1
}
