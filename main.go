package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"iot_device_simulation/internal/cmd"
	_ "iot_device_simulation/internal/logic"
	_ "iot_device_simulation/internal/packed"
	_ "iot_device_simulation/utility"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
