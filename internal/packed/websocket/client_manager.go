package websocket

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
	"sync"
)

type ClientManager struct {
	Clients         map[*Client]bool
	ClientsLock     sync.RWMutex
	Register        chan *Client
	Unregister      chan *Client
	Broadcast       chan *WResponse       // 广播
	ClientBroadcast chan *ClientWResponse // 向某个客户端发送消息
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		Register:        make(chan *Client, 10),
		Unregister:      make(chan *Client, 10),
		Broadcast:       make(chan *WResponse, 1000),
		ClientBroadcast: make(chan *ClientWResponse, 1000),
		Clients:         make(map[*Client]bool),
	}
}

func (cm *ClientManager) HasClient(client *Client) bool {
	cm.ClientsLock.RLock()
	defer cm.ClientsLock.RUnlock()
	_, ok := cm.Clients[client]
	return ok
}

func (cm *ClientManager) GetAllClients() map[*Client]bool {
	clients := make(map[*Client]bool)
	cm.ClientRange(func(client *Client, value bool) bool {
		clients[client] = value
		return true
	})
	return clients
}
func (cm *ClientManager) ClientRange(f func(client *Client, value bool) bool) {
	cm.ClientsLock.RLock()
	defer cm.ClientsLock.RUnlock()
	for client, value := range cm.Clients {
		if !f(client, value) {
			return
		}
	}
}
func (cm *ClientManager) ClientAdd(client *Client) {
	cm.ClientsLock.Lock()
	defer cm.ClientsLock.Unlock()
	cm.Clients[client] = true
}
func (cm *ClientManager) ClientDel(client *Client) {
	cm.ClientsLock.Lock()
	defer cm.ClientsLock.Unlock()
	if _, ok := cm.Clients[client]; ok {
		delete(cm.Clients, client)
	}
}
func (cm *ClientManager) ClientLen() int {
	return len(cm.Clients)
}

// EventRegister 注册事件
func (cm *ClientManager) EventRegister(client *Client) {
	cm.ClientAdd(client)
	client.SendMsg(&WResponse{
		Event: "connect",
		Data: g.Map{
			"ID": client.ID,
		},
	})
}

// EventUnRegister 注销事件
func (cm *ClientManager) EventUnRegister(client *Client) {
	cm.ClientDel(client)
	//if _, ok := cm.Clients[client]; ok {
	//
	//}
}

// clearTimeOutClient 清理超时的客户端
func (cm *ClientManager) clearTimeOutClient() {
	curTime := uint64(gtime.Now().Unix())
	clients := cm.GetAllClients()
	for client, _ := range clients {
		if client.isHeatBeatTimeOut(curTime) {
			_ = client.Socket.Close()
			client.close()
			cm.EventUnRegister(client)
		}
	}
}

func (cm *ClientManager) ping(ctx context.Context) {
	// ping 1分钟执行一次
	_, err := gcron.Add(ctx, "0 */1 * * * *", func(ctx context.Context) {
		res := &WResponse{
			Event: Ping,
			Data:  "ping",
		}
		cm.SendToAll(res)
	}, "websocket_ping")
	if err != nil {
		fmt.Println("websocket ping err:", err)
	}
	// 清理超时链接
	_, err = gcron.Add(ctx, "0 */1 * * * *", func(ctx context.Context) {
		cm.clearTimeOutClient()
	}, "websocket_clear_client")
	if err != nil {
		fmt.Println("websocket client clear err:", err)
	}
}

func (cm *ClientManager) start() {
	for {
		select {
		case client := <-cm.Register:
			cm.EventRegister(client)
		case client := <-cm.Unregister:
			cm.EventUnRegister(client)
		case res := <-cm.Broadcast: //广播
			clients := cm.GetAllClients()
			for client, _ := range clients {
				client.SendMsg(res)
			}
		case res := <-cm.ClientBroadcast: // 单播
			clients := cm.GetAllClients()
			for client, _ := range clients {
				if client.ID == res.ID {
					client.SendMsg(res.WResponse)
				}
			}
		}
	}
}

func (cm *ClientManager) SendToClient(id string, wrs *WResponse) {
	res := &ClientWResponse{
		ID:        id,
		WResponse: wrs,
	}
	cm.ClientBroadcast <- res
}

func (cm *ClientManager) SendToAll(res *WResponse) {
	cm.Broadcast <- res
}
