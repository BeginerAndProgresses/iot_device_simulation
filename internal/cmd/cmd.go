package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"iot_device_simulation/internal/controller/device"
	"iot_device_simulation/internal/controller/mqtt_parameter"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/device", func(group *ghttp.RouterGroup) {
					group.Bind(device.DeviceController)
				})
				group.Group("/mqtt", func(group *ghttp.RouterGroup) {
					group.Bind(mqtt_parameter.MqttController)
				})
				group.GET("/swagger", func(req *ghttp.Request) {
					req.Response.WriteTpl("/swagger.html")
				})

			})
			s.Run()
			return nil
		},
	}
)