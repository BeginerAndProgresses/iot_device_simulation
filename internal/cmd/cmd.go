package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"iot_device_simulation/internal/controller/device"
	"iot_device_simulation/internal/controller/mqtt_parameter"
	"iot_device_simulation/internal/controller/topic"
	"iot_device_simulation/internal/controller/user"
	"iot_device_simulation/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().Auth)
				group.Middleware(service.Middleware().CORS)
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.Bind(user.UserController)
				})
				group.Group("/device", func(group *ghttp.RouterGroup) {
					group.Bind(device.DeviceController)
				})
				group.Group("/mqtt", func(group *ghttp.RouterGroup) {
					group.Bind(mqtt_parameter.MqttController)
				})
				group.Group("/topic", func(group *ghttp.RouterGroup) {
					group.Bind(topic.TopicController)
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
