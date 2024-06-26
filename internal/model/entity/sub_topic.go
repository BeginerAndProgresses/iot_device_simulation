// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SubTopic is the golang structure for table sub_topic.
type SubTopic struct {
	Id       int    `json:"id"        ` //
	SubTopic string `json:"sub_topic" ` // 订阅通信topic
	DeviceId uint   `json:"device_id" ` // 设备id
	State    uint   `json:"state"     ` // 状态，0关闭，1开启
	WsParam  string `json:"ws_param"  ` // 开启ws的参数
	TopicId  int    `json:"topic_id"  ` // 通信id
}
