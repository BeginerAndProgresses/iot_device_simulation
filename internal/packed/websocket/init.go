package websocket

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	clientManager = NewClientManager() // 管理者
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// StartWebSocket 启动WebSocket
func StartWebSocket(ctx context.Context) {
	g.Log().Info(ctx, "启动：WebSocket")
	go clientManager.start()
	go clientManager.ping(ctx)
}

// WebSocketHandler 处理WebSocket请求
func WebSocketHandler(req *ghttp.Request) {
	conn, err := upGrader.Upgrade(req.Response.ResponseWriter, req.Request, nil)
	if err != nil {
		fmt.Println("websocket connection error", err)
	}
	curTime := uint64(gtime.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, curTime)
	go client.read()
	go client.write()
	clientManager.Register <- client
}
