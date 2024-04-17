package websocket

import (
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	Error = "error"
	Login = "login"
	Join  = "join"
	Quit  = "quit"
	IsApp = "is_app"
	Ping  = "ping"
)

func ProcessData(client *Client, message []byte) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("处理数据 stop", r)
		}
	}()
	request := &request{}
	err := gconv.Struct(message, request)
	if err != nil {
		fmt.Println("数据解析失败：", err)
		return
	}
	switch request.Event {
	case Ping:
		PingController(client)
		break
	}
}
