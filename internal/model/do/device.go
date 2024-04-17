// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Device is the golang structure of table device for DAO operations like Where/Data.
type Device struct {
	g.Meta          `orm:"table:device, do:true"`
	Id              interface{} //
	PlatForm        interface{} // å¹³å°åç§°
	DeviceName      interface{} // è®¾å¤‡åç§°
	MqttParameterId interface{} //
	State           interface{} // çŠ¶æ€ï¼Œå¦‚æžœæœªå¯åŠ¨ä¸º0ï¼Œå¦‚æžœå¯åŠ¨ä¸º1
	ProductId       interface{} // äº§å“id
	UserId          interface{} // ç”¨æˆ·id
}
