package service

import (
	"context"
	"fmt"
	"iot_device_simulation/internal/model/entity"
)

var localSimulateDevice ISimulateDevice

type ISimulateDevice interface {
	// Insert 插入数据
	Insert(ctx context.Context, transducers entity.SimulateDevice) error
	// Get 获取数据
	Get(ctx context.Context, id int) (simulateDevice entity.SimulateDevice, err error)
	// Start 启动模拟设备
	Start(ctx context.Context, deviceId int) error
	// Stop 停止模拟设备
	Stop(ctx context.Context, deviceId int) error
	// GetByUid 获取用户下的模拟设备
	GetByUid(ctx context.Context, userid int) (simulateDevice []entity.SimulateDevice, err error)
	// Update 更新数据
	Update(ctx context.Context, transducers entity.SimulateDevice) error
	// Delete 删除数据
	Delete(ctx context.Context, id int) error
}

func SimulateDevice() ISimulateDevice {
	if localTopic == nil {
		fmt.Println("localSimulateDevice未实现或是注册失败")
	}
	return localSimulateDevice
}

func RegisterSimulateDevice(i ISimulateDevice) {
	localSimulateDevice = i
}
