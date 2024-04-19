package websocket

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	CM = NewClientManager() // 管理者
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有源
		return true
	},
}

// StartWebSocket 启动WebSocket
func StartWebSocket(ctx context.Context) {

	//g.Log().Info(ctx, "启动：WebSocket")
	fmt.Println("启动：WebSocket")
	go CM.start()
	go CM.ping(ctx)
}

// WSHandler 处理WebSocket请求
func WSHandler(req *ghttp.Request) {
	conn, err := upGrader.Upgrade(req.Response.ResponseWriter, req.Request, nil)
	if err != nil {
		fmt.Println("websocket connection error", err)
		return
	}
	query := req.URL.Query()
	curTime := uint64(gtime.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, curTime, query.Get("id"))
	go client.read()
	go client.write()
	CM.Register <- client
}
