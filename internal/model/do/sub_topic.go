// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SubTopic is the golang structure of table sub_topic for DAO operations like Where/Data.
type SubTopic struct {
	g.Meta   `orm:"table:sub_topic, do:true"`
	Id       interface{} //
	SubTopic interface{} // 订阅通信topic
	DeviceId interface{} // 设备id
	State    interface{} // 状态，0关闭，1开启
	Url      interface{} // 开启ws通信的路径，以device_id和topic_id拼成
	TopicId  interface{} // 通信id
}
