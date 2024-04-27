package conn

import (
	"errors"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Connect struct {
	PlatForm string
	DeviceId int
	client   mqtt.Client
	option   *mqtt.ClientOptions
}

func NewConnect(platform string, options *mqtt.ClientOptions, deviceId int) *Connect {
	return &Connect{
		PlatForm: platform,
		option:   options,
		DeviceId: deviceId,
	}
}

func (c *Connect) Connect() (err error) {
	if c.PlatForm != "阿里云" {
		c.option.KeepAlive = 30000
	}
	c.client = mqtt.NewClient(c.option)
	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (c *Connect) Publish(topic, json string) (err error) {
	if !c.client.IsConnected() {
		return errors.New("链接未建立")
	}
	token := c.client.Publish(topic, 0, false, json)
	token.Wait()
	err = token.Error()
	return
}

// Subscribe订阅
func (c *Connect) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) (err error) {
	if !c.client.IsConnected() {
		return errors.New("链接未建立")
	}
	token := c.client.Subscribe(topic, qos, callback)
	token.Wait()
	err = token.Error()
	return
}

// Unsubscribe取消订阅
func (c *Connect) Unsubscribe(topic string) (err error) {
	if !c.client.IsConnected() {
		return errors.New("链接未建立")
	}
	token := c.client.Unsubscribe(topic)
	token.Wait()
	err = token.Error()
	return
}

func (c *Connect) Close() {
	c.client.Disconnect(250)
}
