// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Device is the golang structure for table device.
type Device struct {
	Id              int    `json:"id"                ` //
	PlatForm        string `json:"plat_form"         ` // 平台名称
	DeviceName      string `json:"device_name"       ` // 设备名称
	MqttParameterId uint   `json:"mqtt_parameter_id" ` //
	State           uint   `json:"state"             ` // 状态，如果未启动为0，如果启动为1
}
