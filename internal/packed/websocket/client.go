package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"runtime/debug"
)

var (
	heatBeatExpireTime uint64 = 6 * 60
)

type Client struct {
	Addr          string
	ID            string
	Socket        *websocket.Conn
	Send          chan *WResponse
	SendClosed    bool
	FirstTime     uint64 // 首次连接事件
	HeartbeatTime uint64 // 用户上次心跳时间
}

func NewClient(addr string, socket *websocket.Conn, firstTime uint64, id string) *Client {
	return &Client{
		Addr:          addr,
		ID:            id,
		Socket:        socket,
		Send:          make(chan *WResponse, 1000),
		SendClosed:    false,
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
	}
}

func (c *Client) read() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)
		}
	}()
	defer func() {
		c.close()
	}()
	for {
		_, msg, err := c.Socket.ReadMessage()
		if err != nil {
			return
		}
		//	处理信息函数
		fmt.Println("read", c.ID, string(msg))
		ProcessData(c, msg)
	}
}

func (c *Client) write() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)
		}
	}()
	defer func() {
		//	关闭管道
		CM.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				return
			}
			_ = c.Socket.WriteJSON(msg)
		}
	}
}

func (c *Client) SendMsg(msg *WResponse) {
	if c == nil || c.SendClosed {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("send msg stop", string(debug.Stack()), r)
		}
	}()

	c.Send <- msg
}

// updateHeatBeatTime 更新心跳时间
func (c *Client) updateHeatBeatTime(curTime uint64) {
	c.HeartbeatTime = curTime
}

func (c *Client) isHeatBeatTimeOut(curTime uint64) bool {
	if c.HeartbeatTime+heatBeatExpireTime < curTime {
		return true
	}
	return false
}

// 关闭客户端
func (c *Client) close() {
	if c.SendClosed {
		return
	}
	c.SendClosed = true
	close(c.Send)
}
