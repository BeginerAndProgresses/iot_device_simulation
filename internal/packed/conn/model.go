package conn

//	{
//	   "method":"report",
//	   "clientToken":"123",
//	   "timestamp":1628646783,
//	   "params":{
//	       "power_switch":1,
//	       "color":1,
//	       "brightness":32
//	      }
//	}
//
// TencentSendInfo 腾讯云信息格式
type TencentSendInfo struct {
	Method      string      `json:"method"`
	ClientToken string      `json:"clientToken"`
	Params      interface{} `json:"params"`
}

//	{
//	   "id": "123",
//	   "version": "1.0",
//	   "params": {
//	       "temperature": "30.5"
//	   },
//	   "method": "thing.service.property.set"
//	}
//
// AliSendInfo 阿里云信息格式
type AliSendInfo struct {
	Id      string      `json:"id"`
	Version string      `json:"version"`
	Params  interface{} `json:"params"`
	Method  string      `json:"method"`
}

// HuaweiSendInfo 华为云信息格式
type HuaweiSendInfo struct {
	Services []struct {
		ServiceId  string      `json:"service_id"`
		Properties interface{} `json:"properties"`
	} `json:"services"`
}

type SendInfo struct {
	ConnId int
	Topic  string
	Json   string
}

func NewSendInfo(connId int, topic, json string) *SendInfo {
	return &SendInfo{
		ConnId: connId,
		Topic:  topic,
		Json:   json,
	}
}
