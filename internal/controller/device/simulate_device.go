package device

import (
	"context"
	"encoding/json"
	"iot_device_simulation/api/device"
	"iot_device_simulation/internal/controller/mqtt_parameter"
	"iot_device_simulation/internal/model/entity"
	"iot_device_simulation/internal/service"
)

func (c *cDev) AddSimulateDevice(ctx context.Context, req *device.AddSimulateDeviceReq) (res *device.AddSimulateDeviceRes, err error) {
	var mqtt = mqtt_parameter.MqttOption{
		ServiceAddress: req.ServiceAddress,
		Port:           req.Port,
		UserName:       req.UserName,
		Password:       req.Password,
		ClientId:       req.ClientId,
	}
	marshal, err := json.Marshal(&mqtt)
	if err != nil {
		return
	}
	res = &device.AddSimulateDeviceRes{}
	err = service.SimulateDevice().Insert(ctx, entity.SimulateDevice{
		UserId:     req.UserId,
		PlatForm:   req.Platform,
		DeviceId:   req.DeviceId,
		ProductId:  req.ProductId,
		DeviceName: req.DeviceName,
		Topic:      req.Topic,
		Intervals:  uint(req.Interval),
		State:      0,
		MqttOption: string(marshal),
	})
	if err != nil {
		res.Code = 0
		return
	}
	res.Code = 1
	return
}

func (c *cDev) StartSimulateDevice(ctx context.Context, req *device.ConnSimulateDeviceReq) (res *device.ConnSimulateDeviceRes, err error) {
	err = service.SimulateDevice().Start(ctx, req.DeviceId)
	res = &device.ConnSimulateDeviceRes{}
	if err != nil {
		res.State = 0
		return
	}
	res.State = 1
	return
}

func (c *cDev) StopSimulateDevice(ctx context.Context, req *device.DisConnSimulateDeviceReq) (res *device.DisConnSimulateDeviceRes, err error) {
	// 关闭所有的传感器
	err = service.SimulateDevice().Stop(ctx, req.DeviceId)
	res = &device.DisConnSimulateDeviceRes{}
	if err != nil {
		res.State = 0
		return
	}
	res.State = 1
	return
}

func (c *cDev) GetSimulateDevice(ctx context.Context, req *device.GetSimulateDeviceByUidReq) (res *device.GetSimulateDeviceByUidRes, err error) {
	simulateDevice, err := service.SimulateDevice().GetByUid(ctx, req.UserId)
	res = &device.GetSimulateDeviceByUidRes{SimulateDevice: simulateDevice}
	if err != nil {
		res.Code = 0
		return
	}
	res.Code = 1
	return
}

func (c *cDev) UpdateSimulateDevice(ctx context.Context, req *device.EditSimulateDeviceReq) (res *device.EditSimulateDeviceRes, err error) {
	var mqtt = mqtt_parameter.MqttOption{
		ServiceAddress: req.ServiceAddress,
		Port:           req.Port,
		UserName:       req.UserName,
		Password:       req.Password,
		ClientId:       req.ClientId,
	}
	marshal, err := json.Marshal(&mqtt)
	if err != nil {
		return
	}
	err = service.SimulateDevice().Update(ctx, entity.SimulateDevice{
		Id:         req.Id,
		DeviceId:   req.DeviceId,
		Intervals:  uint(req.Interval),
		MqttOption: string(marshal),
	})
	return
}

func (c *cDev) DeleteSimulateDevice(ctx context.Context, req *device.DeleteSimulateDeviceReq) (res *device.DeleteSimulateDeviceRes, err error) {
	err = service.SimulateDevice().Delete(ctx, req.Id)
	res = &device.DeleteSimulateDeviceRes{}
	if err != nil {
		res.Code = 0
		return
	}
	res.Code = 1
	return
}
