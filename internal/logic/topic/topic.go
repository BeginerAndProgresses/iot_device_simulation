package topic

import (
	"context"
	"iot_device_simulation/internal/dao"
	"iot_device_simulation/internal/model/do"
	"iot_device_simulation/internal/model/entity"
	"iot_device_simulation/internal/service"
)

func init() {
	service.RegisterTopic(New())
}

func New() *iTopic {
	return &iTopic{}
}

type iTopic struct {
}

// Get 获取单个Topic
func (i *iTopic) Get(ctx context.Context, id int) (topic entity.Topic, err error) {
	err = dao.Topic.Ctx(ctx).Where("id", id).Scan(&topic)
	return
}

// Insert 插入一个设备
func (i *iTopic) Insert(ctx context.Context, topic do.Topic) (id int, err error) {
	result, err := dao.Topic.Ctx(ctx).Data(topic).Insert()
	insertId, err := result.LastInsertId()
	id = int(insertId)
	return
}

// Update 修改一个设备信息
func (i *iTopic) Update(ctx context.Context, topic do.Topic) (id int, err error) {
	result, err := dao.Topic.Ctx(ctx).Where("id", id).Data(topic).OmitEmptyData().Update()
	insertId, err := result.LastInsertId()
	id = int(insertId)
	return
}

// Delete 删除一个设备
func (i *iTopic) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.Topic.Ctx(ctx).Where("id", id).Delete()
	return
}
