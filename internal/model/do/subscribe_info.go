// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SubscribeInfo is the golang structure of table subscribe_info for DAO operations like Where/Data.
type SubscribeInfo struct {
	g.Meta   `orm:"table:subscribe_info, do:true"`
	Id       interface{} //
	SubName  interface{} //
	Topic    interface{} //
	Info     interface{} // 返回的信息
	DeviceId interface{} //
}
