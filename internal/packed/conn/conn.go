package conn

import (
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"iot_device_simulation/internal/model/entity"
	"sync"
	"time"
)

var (
	connChannel = struct {
		Chans map[int]chan int
		sync.Mutex
	}{}
	connServer = make(map[int]mqtt.Client)
)

func init() {
	connChannel.Chans = make(map[int]chan int)
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

// Conn 普通连接
func Conn(device_id int, parameter *entity.MqttParameter) (err error) {
	fmt.Printf("parameter:", *parameter)
	if client, ok := connServer[device_id]; ok {
		if !client.IsConnected() {
			if token := connServer[device_id].Connect(); token.Wait() && token.Error() != nil {
				return token.Error()
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
			opts.KeepAlive = 30000
		}
		client = mqtt.NewClient(opts)
		connServer[device_id] = client
		if token := connServer[device_id].Connect(); token.Wait() && token.Error() != nil {
			return token.Error()
		}
	}
	//connChannel[device_id] = make(chan int)

	return nil
}

// ConcConn 并发连接
func ConcConn(device_id int, parameter *entity.MqttParameter) (err error) {
	connChannel.Lock()
	connChannel.Chans[device_id] = make(chan int)
	connChannel.Unlock()
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
		opts.KeepAlive = 30000
	}
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	<-connChannel.Chans[device_id]
	client.Disconnect(250)
	return
}

func DisConn(device_id int) (err error) {
	if client, ok := connServer[device_id]; !ok {
		fmt.Printf("不存在连接 device_id", device_id)
	} else {
		client.Disconnect(250)
	}
	return
}

// DisConcConn 并发关闭连接
func DisConcConn(device_id int) (err error) {
	connChannel.Lock()
	connChannel.Chans[device_id] <- 0
	connChannel.Unlock()
	return
}

// Publish 推送数据
func Publish(device_id int, topic, json string) (err error) {
	fmt.Printf("---------------")
	fmt.Printf("json", json, "topic", topic)
	if client, ok := connServer[device_id]; !ok {
		fmt.Printf("不存在连接 device_id", device_id)
		return errors.New(fmt.Sprintf("不存在连接 device_id:%d", device_id))
	} else {
		if client.IsConnected() {
			fmt.Printf("尝试发送")
			token := client.Publish(topic, 0, false, json)
			token.Wait()
			time.Sleep(time.Second)
		} else {
			fmt.Printf("发送失败")
		}
	}
	return
}

// Subscribe 订阅topic
func Subscribe(device_id int, topic string) (err error) {
	if client, ok := connServer[device_id]; !ok {
		fmt.Printf("不存在连接 device_id", device_id)
		return errors.New(fmt.Sprintf("不存在连接 device_id:%d", device_id))
	} else {
		if client.IsConnected() {
			fmt.Printf("尝试订阅")
			token := client.Subscribe(topic, 1, nil)
			token.Wait()
		} else {
			fmt.Printf("订阅失败")
		}
	}
	return
}
