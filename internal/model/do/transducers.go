// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Transducers is the golang structure of table transducers for DAO operations like Where/Data.
type Transducers struct {
	g.Meta          `orm:"table:transducers, do:true"`
	Id              interface{} //
	UserId          interface{} // 用户id
	DeviceId        interface{} // 设备id
	TransducersType interface{} // 传感器类型
	Identifier      interface{} // 标识符
	Option          interface{} // 配置，传入json
}
