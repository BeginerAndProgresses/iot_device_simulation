package websocket

import (
	"github.com/gogf/gf/v2/os/gtime"
)

func PingController(client *Client) {
	currentTime := uint64(gtime.Now().Unix())
	client.updateHeatBeatTime(currentTime)
}
