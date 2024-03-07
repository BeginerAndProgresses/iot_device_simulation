package conn

import (
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"iot_device_simulation/internal/model/entity"
)

var (
	//connChannel = make(map[int]chan int)
	connServer = make(map[int]mqtt.Client)
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func Conn(device_id int, parameter *entity.MqttParameter) (err error) {
	if client, ok := connServer[device_id]; ok {
		if !client.IsConnected() {
			if token := connServer[device_id].Connect(); token.Wait() && token.Error() != nil {
				return errors.New("连接失败")
			}
		} else {
			return
		}
	} else {
		fmt.Printf("开始创建连接 device_id", device_id)
		var broker = parameter.ServerAddress
		var port = parameter.Port
		opts := mqtt.NewClientOptions()
		{
			opts.AddBroker(fmt.Sprintf("mqtt://%s:%d", broker, port))
			opts.SetClientID(parameter.ClientId)
			opts.SetUsername(parameter.Username)
			opts.SetPassword(parameter.Password)
			opts.SetDefaultPublishHandler(messagePubHandler)
			opts.OnConnect = connectHandler
			opts.OnConnectionLost = connectLostHandler
		}
		client = mqtt.NewClient(opts)
		connServer[device_id] = client
		if token := connServer[device_id].Connect(); token.Wait() && token.Error() != nil {
			return errors.New("连接失败")
		}
	}
	//connChannel[device_id] = make(chan int)

	return nil
}

func DisConn(device_id int) (err error) {
	if client, ok := connServer[device_id]; !ok {
		fmt.Printf("不存在连接 device_id", device_id)
	} else {
		client.Disconnect(250)
	}
	return
}
