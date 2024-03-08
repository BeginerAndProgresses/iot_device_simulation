package service

import (
	"context"
	"fmt"
	"iot_device_simulation/internal/model/do"
	"iot_device_simulation/internal/model/entity"
)

type IDevice interface {
	// Get 获取单个设备
	Get(ctx context.Context, id int) (device entity.Device, err error)
	// Insert 插入一个设备
	Insert(ctx context.Context, device do.Device) (id int, err error)
	// Update 修改一个设备信息
	Update(ctx context.Context, device do.Device) (id int, err error)
	// Delete 删除一个设备
	Delete(ctx context.Context, id int) (err error)
	// ConnMqtt 使用Mqtt参数完成设备上线
	ConnMqtt(ctx context.Context, deviceId int) (id, state int, err error)
	// DisConnMqtt 设备下线
	DisConnMqtt(ctx context.Context, deviceId int) (id, state int, err error)
	// InfoPost 属性上报
	InfoPost(ctx context.Context, deviceId, topicId int, json string) (err error)
}

var localDevice IDevice

func Device() IDevice {
	if localDevice == nil {
		fmt.Println("localDevice未实现或是注册失败")
	}
	return localDevice
}

func RegisterDevice(i IDevice) {
	localDevice = i
}
