package websocket

import "github.com/gogf/gf/v2/frame/g"

type request struct {
	Event string `json:"event"` //事件名称
	Data  g.Map  `json:"data"`  //数据
}

type WResponse struct {
	Event string      `json:"event"` // 事件名称
	Data  interface{} `json:"data"`  // 数据
}

type ClientWResponse struct {
	ID        string     `json:"id"`
	WResponse *WResponse `json:"wResponse"`
}
