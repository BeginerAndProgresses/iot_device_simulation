package websocket

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PingReq struct {
	g.Meta `path:"/ping" method:"get" summary:"You first hello api"`
}

type PingRes struct {
}

//type ConnectionReq struct {
//	g.Meta  `path:"/socket" method:"all" summary:"You first hello api"`
//	ID      string  `p:"id" v:"required|integer|min:1#id不能为空|id只能是整数|最小值不应小于1" dc:"id"`
//	Request Request `p:"request" v:"required#request不能为空" dc:"请求"`
//}
//
//type Request struct {
//	Event string `p:"event"`
//	Data  g.Map  `p:"data"`
//}
//
//type ConnectionRes struct {
//	Code int `json:"code"`
//}
