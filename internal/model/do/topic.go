// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Topic is the golang structure of table topic for DAO operations like Where/Data.
type Topic struct {
	g.Meta           `orm:"table:topic, do:true"`
	Id               interface{} // id
	PlatForm         interface{} // 平台名
	Topic            interface{} // 通信topic
	FunctionDescribe interface{} // 功能描述
}
