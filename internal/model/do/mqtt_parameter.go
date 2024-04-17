// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MqttParameter is the golang structure of table mqtt_parameter for DAO operations like Where/Data.
type MqttParameter struct {
	g.Meta        `orm:"table:mqtt_parameter, do:true"`
	Id            interface{} //
	ClientId      interface{} // Client ID
	Port          interface{} // ç«¯å£å·
	ServerAddress interface{} // æœåŠ¡å™¨åœ°å€
	Username      interface{} // username
	Password      interface{} // password
	UserId        interface{} // ç”¨æˆ·id
}
