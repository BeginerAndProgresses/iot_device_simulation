// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SimulateDevice is the golang structure for table simulate_device.
type SimulateDevice struct {
	Id         int    `json:"id"          ` //
	PlatForm   string `json:"plat_form"   ` // 平台
	DeviceName string `json:"device_name" ` // 设备名
	DeviceId   string `json:"device_id"   ` // 设备id
	State      uint   `json:"state"       ` // 状态
	ProductId  string `json:"product_id"  ` // 产品id
	UserId     int    `json:"user_id"     ` // 用户id
	Intervals  uint   `json:"intervals"   ` // 时间间隔，默认20秒
	MqttOption string `json:"mqtt_option" ` // mqtt连接参数
	Topic      string `json:"topic"       ` // 通信主题
}
