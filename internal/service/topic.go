package service

import (
	"context"
	"fmt"
	"iot_device_simulation/internal/model/do"
	"iot_device_simulation/internal/model/entity"
)

var localTopic ITopic

type ITopic interface {
	// Get 获取单个topic
	Get(ctx context.Context, id int) (topic entity.Topic, err error)
	// Insert 插入一个topic
	Insert(ctx context.Context, topic do.Topic) (id int, err error)
	// Update 修改一个topic信息
	Update(ctx context.Context, topic do.Topic) (id int, err error)
	// Delete 删除一个topic
	Delete(ctx context.Context, id int) (err error)
	// GetAllUpTopics 获取所有上传topic
	GetAllUpTopics(ctx context.Context, userId int) (topics []entity.Topic, err error)
	// GetAllDownTopics 获取所有下拉topic
	GetAllDownTopics(ctx context.Context, userId int) (topics []entity.Topic, err error)
	// GetAllUpByDeviceIdTopics 根据设备id获取平台所有上传topic
	GetAllUpByDeviceIdTopics(ctx context.Context, userId, deviceId int) (topics []entity.Topic, err error)
	// GetAllDownByDeviceIdTopics 根据设备id获取平台所有下拉topic
	GetAllDownByDeviceIdTopics(ctx context.Context, userId, deviceId int) (topics []entity.Topic, err error)
}

func Topic() ITopic {
	if localTopic == nil {
		fmt.Println("localDevice未实现或是注册失败")
	}
	return localTopic
}

func RegisterTopic(i ITopic) {
	localTopic = i
}
