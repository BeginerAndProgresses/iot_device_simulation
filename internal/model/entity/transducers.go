// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Transducers is the golang structure for table transducers.
type Transducers struct {
	Id              int    `json:"id"               ` //
	UserId          int    `json:"user_id"          ` // 用户id
	DeviceId        uint   `json:"device_id"        ` // 设备id
	TransducersType string `json:"transducers_type" ` // 传感器类型
	Identifier      string `json:"identifier"       ` // 标识符
	Option          string `json:"option"           ` // 配置，传入json
}
