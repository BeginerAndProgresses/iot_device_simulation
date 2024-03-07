package device

import (
	"context"
	"fmt"
	"iot_device_simulation/api/device"
	"iot_device_simulation/internal/model/do"
	"iot_device_simulation/internal/service"
)

// DeviceController 用于注册路由
var DeviceController = &cDev{}

type cDev struct {
}

func (c *cDev) Add(ctx context.Context, req *device.AddReq) (res *device.AddRes, err error) {
	id, err := service.Device().Insert(ctx, do.Device{
		DeviceName: req.DeviceName,
		PlatForm:   req.Platform,
	})
	res = &device.AddRes{Id: id}
	return
}

func (c *cDev) Search(ctx context.Context, req *device.SearchReq) (res *device.SearchRes, err error) {
	get, err := service.Device().Get(ctx, req.Id)
	{
		res = &device.SearchRes{}
		res.Device.Id = get.Id
		res.Device.PlatForm = get.PlatForm
		res.Device.DeviceName = get.DeviceName
		res.Device.MqttParameterId = get.MqttParameterId
	}
	fmt.Println(*res)
	return
}

func (c *cDev) Del(ctx context.Context, req *device.DelReq) (res *device.DelRes, err error) {
	err = service.Device().Delete(ctx, req.Id)
	res = &device.DelRes{}
	if err != nil {
		res.Id = 0
	}
	res.Id = req.Id
	return
}

func (c *cDev) Edit(ctx context.Context, req *device.EditReq) (res *device.EditRes, err error) {
	id, err := service.Device().Update(ctx, do.Device{
		Id:              req.Id,
		DeviceName:      req.DeviceName,
		PlatForm:        req.Platform,
		MqttParameterId: req.MqttParameterId,
	})
	res = &device.EditRes{}
	res.Id = id
	return
}

func (c *cDev) Conn(ctx context.Context, req *device.ConnReq) (res *device.ConnRes, err error) {
	id, state, err := service.Device().ConnMqtt(ctx, req.Id)
	res = &device.ConnRes{Id: id, State: state}
	return
}

func (c *cDev) DisConn(ctx context.Context, req *device.DisConnReq) (res *device.DisConnRes, err error) {
	id, state, err := service.Device().DisConnMqtt(ctx, req.Id)
	res = &device.DisConnRes{Id: id, State: state}
	return
}
