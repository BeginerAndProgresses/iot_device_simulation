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
	// GetByPlatform 根据平台名获取设备信息
	GetByPlatform(ctx context.Context, platform string, userid int) (devices []entity.Device, err error)
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
	// InfoPost 信息上报
	InfoPost(ctx context.Context, userId, deviceId, topicId int, json string) (err error)
	// TopicSub 订阅topic
	TopicSub(ctx context.Context, deviceId, topicId int) (err error)
	// AddSubTopic 添加订阅topic
	AddSubTopic(ctx context.Context, subTopic do.SubTopic) error
	// GetByUID 获取用户设备
	GetByUID(ctx context.Context, userid int) (TencentDevice, HuaweiDevice, AliyunDevice []entity.Device, err error)
	// GetChatDataInfo 获取图表数据
	GetChatDataInfo(ctx context.Context, userid int, days int) (times []string, lineData, barOnlineData, barOffOnlineData []int, err error)
	// GetSubTopic 获取设备订阅的topic
	GetSubTopic(ctx context.Context, deviceId int) (topics []entity.SubTopic, err error)
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
