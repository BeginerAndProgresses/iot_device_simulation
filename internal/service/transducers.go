package service

import (
	"context"
	"fmt"
	"iot_device_simulation/internal/model/entity"
)

var localTransducers ITransducers

type ITransducers interface {
	// Insert 添加传感器
	Insert(ctx context.Context, transducers entity.Transducers) error
	// StopAllTransducers 暂停所有传感器
	StopAllTransducers(ctx context.Context, deviceId int) error
	// StartAllTransducers 启动所有传感器
	StartAllTransducers(ctx context.Context, deviceId int, interval int64) error
	// GetAllTransducersByDeviceId 获取设备所有传感器
	GetAllTransducersByDeviceId(ctx context.Context, deviceId int) ([]entity.Transducers, error)
	// GetAllTransducersByUidPage 获取用户所有传感器
	GetAllTransducersByUidPage(ctx context.Context, uid int, page int, size int) (transducers []entity.Transducers, err error)
	// GetAllTransducersCountByUid 获取用户所有传感器总数
	GetAllTransducersCountByUid(ctx context.Context, uid int) (count int, err error)
	// Get 获取单个传感器
	Get(ctx context.Context, id int) (entity.Transducers, error)
	// Update 更新传感器
	Update(ctx context.Context, transducers entity.Transducers) error
	// Delete 删除传感器
	Delete(ctx context.Context, id int) error
	// SetDeviceId 设置设备id
	SetDeviceId(ctx context.Context, id int, deviceId int) error
	// GetAllTransducersByUidButDeviceId 查询未被使用的传感器
	GetAllTransducersByUidButDeviceId(ctx context.Context, uid int) ([]entity.Transducers, error)
}

func Transducers() ITransducers {
	if localTopic == nil {
		fmt.Println("localTransducers未实现或是注册失败")
	}
	return localTransducers
}

func RegisterTransducers(i ITransducers) {
	localTransducers = i
}
