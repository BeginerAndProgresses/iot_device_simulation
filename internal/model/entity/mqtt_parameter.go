// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// MqttParameter is the golang structure for table mqtt_parameter.
type MqttParameter struct {
	Id            int    `json:"id"             ` //
	ClientId      string `json:"client_id"      ` // Client ID
	Port          int    `json:"port"           ` // 端口号
	ServerAddress string `json:"server_address" ` // 服务器地址
	Username      string `json:"username"       ` // username
	Password      string `json:"password"       ` // password
}
