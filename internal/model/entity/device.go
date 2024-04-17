// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Device is the golang structure for table device.
type Device struct {
	Id              int    `json:"id"                ` //
	PlatForm        string `json:"plat_form"         ` // å¹³å°åç§°
	DeviceName      string `json:"device_name"       ` // è®¾å¤‡åç§°
	MqttParameterId uint   `json:"mqtt_parameter_id" ` //
	State           uint   `json:"state"             ` // çŠ¶æ€ï¼Œå¦‚æžœæœªå¯åŠ¨ä¸º0ï¼Œå¦‚æžœå¯åŠ¨ä¸º1
	ProductId       string `json:"product_id"        ` // äº§å“id
	UserId          int    `json:"user_id"           ` // ç”¨æˆ·id
}
