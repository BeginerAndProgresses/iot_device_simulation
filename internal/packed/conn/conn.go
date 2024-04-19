package conn

import (
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogf/gf/v2/errors/gerror"
	"iot_device_simulation/internal/model/entity"
	"iot_device_simulation/internal/packed/util"
	"sync"
	"time"
)

var (
	connChannel = struct {
		Chans                map[int]chan int
		PublishTopicChan     map[int]chan string
		PublishJsonChan      map[int]chan string
		SubscribeTopicChan   map[int]chan string
		SubscribeFuncChan    map[int]chan func(int, int, string) func(mqtt.Client, mqtt.Message)
		SubscribeTopicIdChan map[int]chan int
		sync.Mutex
	}{}
	connServer = make(map[int]mqtt.Client)
)

func init() {
	connChannel.Chans = make(map[int]chan int)
	connChannel.PublishTopicChan = make(map[int]chan string)
	connChannel.PublishJsonChan = make(map[int]chan string)
	connChannel.SubscribeTopicChan = make(map[int]chan string)
	connChannel.SubscribeFuncChan = make(map[int]chan func(int, int, string) func(mqtt.Client, mqtt.Message))
	connChannel.SubscribeTopicIdChan = make(map[int]chan int)
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
func ConcConn(device entity.Device, parameter *entity.MqttParameter) (err error) {
	connChannel.Lock()
	connChannel.Chans[device.Id] = make(chan int)
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
		if device.PlatForm != "华为云" {
			opts.KeepAlive = 60000 // 腾讯云 0-90秒 阿里云60秒起 华为云标的默认60秒，但是如果配置将会认证失败
		}
	}
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("Error connecting device_id:%d,Error:%v\n", device.Id, token.Error())
		return token.Error()
	}
	// connChannel.Chans[device_id] 通道中传入不同值进行不同操作，0停止，1发送，2订阅
	for client.IsConnected() {
		select {
		case i := <-connChannel.Chans[device.Id]:
			switch i {
			case 0:
				client.Disconnect(250)
				return
			case 1:
				topic := <-connChannel.PublishTopicChan[device.Id]
				json := <-connChannel.PublishJsonChan[device.Id]
				token := client.Publish(topic, 0, false, json)
				token.Wait()
				close(connChannel.PublishTopicChan[device.Id])
				close(connChannel.PublishJsonChan[device.Id])
			case 2:
				topicId := <-connChannel.SubscribeTopicIdChan[device.Id]
				topic := <-connChannel.SubscribeTopicChan[device.Id]
				subFunc := <-connChannel.SubscribeFuncChan[device.Id]
				token := client.Subscribe(topic, 1, subFunc(device.Id, topicId, topic))
				token.Wait()
				close(connChannel.SubscribeTopicChan[device.Id])
				close(connChannel.SubscribeFuncChan[device.Id])
			}
		}
	}
	fmt.Printf("尝试关闭连接 device_id", device.Id)
	connChannel.Lock()
	close(connChannel.Chans[device.Id])
	connChannel.Unlock()
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
	// 判断是否存在连接，根据通道判断
	if util.IsChanIntClose(connChannel.Chans[device_id]) {
		err = gerror.New("连接已关闭")
		return
	}
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

// ConcPublish 推送数据
func ConcPublish(device_id int, topic, json string) (err error) {
	if util.IsChanIntClose(connChannel.Chans[device_id]) {
		err = gerror.New("连接已关闭")
		return
	}
	connChannel.Lock()
	connChannel.PublishTopicChan[device_id] = make(chan string)
	connChannel.PublishJsonChan[device_id] = make(chan string)
	connChannel.Unlock()
	connChannel.Chans[device_id] <- 1
	connChannel.PublishTopicChan[device_id] <- topic
	connChannel.PublishJsonChan[device_id] <- json
	return
}

// ConnSubscribe 并发订阅topic
func ConnSubscribe(device_id, topic_id int, topic string, subBack func(int, int, string) func(mqtt.Client, mqtt.Message)) (err error) {
	if util.IsChanIntClose(connChannel.Chans[device_id]) {
		err = gerror.New("连接已关闭")
		return
	}
	connChannel.Lock()
	connChannel.SubscribeTopicChan[device_id] = make(chan string)
	connChannel.SubscribeFuncChan[device_id] = make(chan func(int, int, string) func(mqtt.Client, mqtt.Message))
	connChannel.SubscribeTopicIdChan[device_id] = make(chan int)
	connChannel.Unlock()
	connChannel.Chans[device_id] <- 2
	connChannel.SubscribeTopicIdChan[device_id] <- topic_id
	connChannel.SubscribeTopicChan[device_id] <- topic
	connChannel.SubscribeFuncChan[device_id] <- subBack
	return
}
