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
	PlatForm         interface{} // å¹³å°å
	Topic            interface{} // é€šä¿¡topic
	FunctionDescribe interface{} // åŠŸèƒ½æè¿°
	UserId           interface{} // ç”¨æˆ·id
	TType            interface{} // 订阅为1，上报为0
}
