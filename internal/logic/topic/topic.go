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
	if err != nil {
		return
	}
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

func (i *iTopic) GetAllUpTopics(ctx context.Context, userId int) (topics []entity.Topic, err error) {
	err = dao.Topic.Ctx(ctx).Where("user_id", userId).Where(dao.Topic.Columns().TType, 1).Scan(&topics)
	return
}

func (i *iTopic) GetAllDownTopics(ctx context.Context, userId int) (topics []entity.Topic, err error) {
	err = dao.Topic.Ctx(ctx).Where("user_id", userId).Where(dao.Topic.Columns().TType, 0).Scan(&topics)
	return
}

func (i *iTopic) GetAllUpByDeviceIdTopics(ctx context.Context, userId, deviceId int) (topics []entity.Topic, err error) {
	device, err := service.Device().Get(ctx, deviceId)
	if err != nil {
		return []entity.Topic{}, err
	}
	err = dao.Topic.Ctx(ctx).Where("user_id", userId).Where(dao.Topic.Columns().PlatForm, device.PlatForm).Where(dao.Topic.Columns().TType, 0).Scan(&topics)
	return
}

func (i *iTopic) GetAllDownByDeviceIdTopics(ctx context.Context, userId, deviceId int) (topics []entity.Topic, err error) {
	device, err := service.Device().Get(ctx, deviceId)
	if err != nil {
		return []entity.Topic{}, err
	}
	err = dao.Topic.Ctx(ctx).Where("user_id", userId).Where(dao.Topic.Columns().PlatForm, device.PlatForm).Where(dao.Topic.Columns().TType, 1).Scan(&topics)
	return
}

func (i *iTopic) GetAll(ctx context.Context, userId int) (topics []entity.Topic, err error) {
	err = dao.Topic.Ctx(ctx).Where("user_id", userId).Scan(&topics)
	return
}

func (i *iTopic) GetAllByPage(ctx context.Context, userId int, page, size int) (topics []entity.Topic, allSize int, err error) {
	err = dao.Topic.Ctx(ctx).Where("user_id", userId).Page(page, size).Scan(&topics)
	allSize, err = dao.Topic.Ctx(ctx).Where("user_id", userId).Count()
	return
}

func (i *iTopic) GetAllByPageAndMsg(ctx context.Context, userId int, page, size int, msg string) (topics []entity.Topic, allSize int, err error) {
	if msgInPlatFrom(msg) {
		err = dao.Topic.Ctx(ctx).Where("user_id", userId).Where(dao.Topic.Columns().PlatForm, msg).Page(page, size).Scan(&topics)
		allSize, err = dao.Topic.Ctx(ctx).Where("user_id", userId).Where(dao.Topic.Columns().PlatForm, msg).Count()
		return
	}
	if msgInTType(msg) {
		if msg == "上报" {
			err = dao.Topic.Ctx(ctx).Where("user_id", userId).Where(dao.Topic.Columns().TType, 0).Page(page, size).Scan(&topics)
			allSize, err = dao.Topic.Ctx(ctx).Where("user_id", userId).Where(dao.Topic.Columns().TType, 0).Count()
		} else {
			err = dao.Topic.Ctx(ctx).Where("user_id", userId).Where(dao.Topic.Columns().TType, 1).Page(page, size).Scan(&topics)
			allSize, err = dao.Topic.Ctx(ctx).Where("user_id", userId).Where(dao.Topic.Columns().TType, 1).Count()
		}
		return
	}
	err = dao.Topic.Ctx(ctx).Where("user_id", userId).Where(dao.Topic.Columns().Topic, "%"+msg+"%").Page(page, size).Scan(&topics)
	allSize, err = dao.Topic.Ctx(ctx).Where("user_id", userId).Where(dao.Topic.Columns().Topic, "%"+msg+"%").Count()
	return
}

func msgInPlatFrom(msg string) bool {
	return msg == "阿里云" || msg == "腾讯云" || msg == "华为云"
}

func msgInTType(msg string) bool {
	return msg == "上报" || msg == "下行"
}
