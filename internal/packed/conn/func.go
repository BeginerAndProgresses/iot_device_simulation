package conn

import (
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"iot_device_simulation/internal/packed/util"
	"iot_device_simulation/internal/service"
)

func StartMqttServe() error {
	CManager = NewConnectManager()
	go CManager.start()
	fmt.Println("启动mqtt服务成功")
	return nil
}

var (
	connectLostFunc mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("连接丢失: %v", err)
	}
	connectFunc mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("开始连接...")
	}
	pubFunc mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("发送的数据为%s,信息为%s", msg.Payload(), msg.Topic())
	}
)

func CreatAConn(ctx context.Context, deviceId int) {
	device, _ := service.SimulateDevice().Get(ctx, deviceId)
	var mp struct {
		ServiceAddress string `json:"service_address"`
		Port           int    `json:"port"`
		ClientId       string `json:"client_id"`
		UserName       string `json:"user_name"`
		Password       string `json:"password"`
	}
	err := json.Unmarshal([]byte(device.MqttOption), &mp)
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	options := mqtt.NewClientOptions()
	{
		options.AddBroker(fmt.Sprintf("mqtt://%s:%d", mp.ServiceAddress, mp.Port))
		options.SetClientID(mp.ClientId)
		options.SetUsername(mp.UserName)
		options.SetPassword(mp.Password)
		options.SetConnectionLostHandler(connectLostFunc)
		options.SetOnConnectHandler(connectFunc)
		options.SetDefaultPublishHandler(pubFunc)
		options.SetAutoReconnect(true)
	}
	connect := NewConnect(device.PlatForm, options, deviceId)
	err = connect.Connect()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	CManager.ConnectChan <- connect
}

func CloseAConn(deviceId int) {
	CManager.RangeConnect(func(connect *Connect, value bool) bool {
		if connect.DeviceId == deviceId {
			connect.Close()
			CManager.CloseChan <- connect
			return false
		}
		return true
	})
}

func TopicPadding(platform string, topic string, deviceName, productId, deviceId string) string {
	m := make(map[string]string)
	var newTopic string
	switch platform {
	case "阿里云":
		m["deviceName"] = deviceName
		m["productKey"] = productId
		newTopic = util.VariableString2String(topic, m, "${", "}")
	case "腾讯云":
		m["DeviceName"] = deviceName
		m["ProductID"] = productId
		newTopic = util.VariableString2String(topic, m, "{", "}")
	case "OneNet":
	case "华为云":
		m["device_id"] = deviceId
		newTopic = util.VariableString2String(topic, m, "{", "}")
	}
	return newTopic
}
