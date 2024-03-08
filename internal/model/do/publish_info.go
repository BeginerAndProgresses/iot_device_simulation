// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PublishInfo is the golang structure of table publish_info for DAO operations like Where/Data.
type PublishInfo struct {
	g.Meta `orm:"table:publish_info, do:true"`
	Id     interface{} // id
	Json   interface{} // 信息json
	Topic  interface{} // 通信topic
}
