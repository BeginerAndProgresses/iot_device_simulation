package mqtt_parameter

type MqttOption struct {
	UserName       string `json:"user_name"`
	Password       string `json:"password"`
	ClientId       string `json:"client_id"`
	ServiceAddress string `json:"service_address"`
	Port           int    `json:"port"`
}
