package conn

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/util/guid"
	"sync"
)

var (
	CManager *ConnectManager
)

type ConnectManager struct {
	connects      map[*Connect]bool
	connectLock   sync.RWMutex
	ConnectChan   chan *Connect
	CloseChan     chan *Connect
	ConnBroadcast chan *SendInfo
}

func NewConnectManager() *ConnectManager {
	return &ConnectManager{
		connects:      make(map[*Connect]bool),
		ConnectChan:   make(chan *Connect, 100),
		CloseChan:     make(chan *Connect, 100),
		ConnBroadcast: make(chan *SendInfo, 100),
	}
}

func (cm *ConnectManager) HasConnect(connect *Connect) bool {
	cm.connectLock.RLock()
	defer cm.connectLock.RUnlock()
	_, ok := cm.connects[connect]
	return ok
}

func (cm *ConnectManager) AddConnect(connect *Connect) {
	cm.connectLock.Lock()
	defer cm.connectLock.Unlock()
	cm.connects[connect] = true
}

func (cm *ConnectManager) DelConnect(connect *Connect) {
	cm.connectLock.Lock()
	defer cm.connectLock.Unlock()
	delete(cm.connects, connect)
}

func (cm *ConnectManager) RangeConnect(f func(connect *Connect, value bool) bool) {
	cm.connectLock.RLock()
	defer cm.connectLock.RUnlock()
	for connect, value := range cm.connects {
		if !f(connect, value) {
			break
		}
	}
}

func (cm *ConnectManager) PublishToTencent(deviceId int, topic string, msg map[string]interface{}) {
	var tencentSendInfo = new(TencentSendInfo)
	tencentSendInfo.Method = "report"
	tencentSendInfo.ClientToken = guid.S()
	tencentSendInfo.Params = msg
	marshal, err := json.Marshal(tencentSendInfo)
	if err != nil {
		fmt.Println("marshal tencentSendInfo err:", err)
		return
	}
	clear(msg)
	cm.ConnBroadcast <- NewSendInfo(deviceId, topic, string(marshal))
}

func (cm *ConnectManager) PublishToAli(deviceId int, topic string, msg map[string]interface{}) {
	var aliSendInfo = new(AliSendInfo)
	aliSendInfo.Params = msg
	aliSendInfo.Id = guid.S()
	aliSendInfo.Version = "1.0"
	aliSendInfo.Method = "thing.service.property.set"
	marshal, err := json.Marshal(aliSendInfo)
	if err != nil {
		fmt.Println("marshal tencentSendInfo err:", err)
		return
	}
	cm.ConnBroadcast <- NewSendInfo(deviceId, topic, string(marshal))
	clear(msg)
}

func (cm *ConnectManager) PublishToHuawei(deviceId int, topic string, msg map[string]interface{}, HuaWeiDeviceId string) {
	var huaweiSendInfo = new(HuaweiSendInfo)
	huaweiSendInfo.Services = []struct {
		ServiceId  string      `json:"service_id"`
		Properties interface{} `json:"properties"`
	}{
		{
			ServiceId:  HuaWeiDeviceId,
			Properties: msg,
		},
	}
	marshal, err := json.Marshal(huaweiSendInfo)
	if err != nil {
		fmt.Println("marshal huaweiSendInfo err:", err)
		return
	}
	clear(msg)
	cm.ConnBroadcast <- NewSendInfo(deviceId, topic, string(marshal))
}

func (cm *ConnectManager) start() {
	for {
		select {
		case connect := <-cm.ConnectChan:
			cm.AddConnect(connect)
		case sendInfo := <-cm.ConnBroadcast:
			fmt.Println("sendInfo:", sendInfo.ConnId, sendInfo.Topic, sendInfo.Json)
			cm.RangeConnect(func(connect *Connect, value bool) bool {
				if connect.DeviceId == sendInfo.ConnId {
					_ = connect.Publish(sendInfo.Topic, sendInfo.Json)
				}
				return true
			})
		}
	}
}
