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
		ProductId:  req.ProductId,
		UserId:     req.UserId,
	})
	res = &device.AddRes{Id: id}
	return
}

func (c *cDev) Search(ctx context.Context, req *device.SearchReq) (res *device.SearchRes, err error) {
	get, err := service.Device().Get(ctx, req.Id)
	{
		res = &device.SearchRes{}
		res.Device = get
	}
	fmt.Println(*res)
	return
}

func (c *cDev) SearchByPlatform(ctx context.Context, req *device.SearchByPlatformReq) (res *device.SearchByPlatformRes, err error) {
	devices, err := service.Device().GetByPlatform(ctx, req.Platform, req.UserId)
	{
		res = &device.SearchByPlatformRes{}
		res.Devices = devices
	}
	//fmt.Println(*res)
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
		ProductId:       req.ProductId,
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

// TopicPost 上传属性
func (c *cDev) TopicPost(ctx context.Context, req *device.PubReq) (res *device.PubRes, err error) {
	err = service.Device().InfoPost(ctx, req.DeviceId, req.TopicId, req.JsonInfo)
	res = &device.PubRes{}
	if err != nil {
		res.Code = 0
	} else {
		res.Code = 1
	}
	return
}

// TopicSub topic订阅
func (c *cDev) TopicSub(ctx context.Context, req *device.SubReq) (res *device.SubRes, err error) {
	err = service.Device().TopicSub(ctx, req.DeviceId, req.TopicId)
	res = &device.SubRes{}
	if err != nil {
		res.Code = 0
	} else {
		res.Code = 1
	}
	return
}

// TopicSubInfo topic订阅返回信息获取
func (c *cDev) TopicSubInfo(ctx context.Context, req *device.SubInfoReq) (res *device.SubInfoRes, err error) {
	return
}
