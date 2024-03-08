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
