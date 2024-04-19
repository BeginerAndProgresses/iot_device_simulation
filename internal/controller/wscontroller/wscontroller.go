package wscontroller

import (
	"context"
	"iot_device_simulation/api/websocket"
)

var WebSocketController = &cWebSocket{}

type cWebSocket struct {
}

func (c *cWebSocket) Ping(ctx context.Context, req *websocket.PingReq) (res *websocket.PingRes, err error) {
	return
}

//func (c *cWebSocket) Connection(ctx context.Context, req *websocket.ConnectionReq) (res *websocket.ConnectionRes, err error) {
//	return
//}
