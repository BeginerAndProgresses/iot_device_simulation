package topic

import (
	"context"
	"iot_device_simulation/api/topic"
	"iot_device_simulation/internal/model/do"
	"iot_device_simulation/internal/service"
)

var TopicController = &cTopic{}

type cTopic struct {
}

func (c *cTopic) Add(ctx context.Context, req *topic.AddReq) (res *topic.AddRes, err error) {
	id, err := service.Topic().Insert(ctx, do.Topic{
		UserId:           req.UserId,
		PlatForm:         req.PlatForm,
		Topic:            req.Topic,
		FunctionDescribe: req.FunctionDescribe,
		TType:            req.TType,
	})
	res = &topic.AddRes{Id: id}
	return
}

func (c *cTopic) Search(ctx context.Context, req *topic.SearchReq) (res *topic.SearchRes, err error) {
	get, err := service.Topic().Get(ctx, req.Id)
	res = &topic.SearchRes{}
	if err != nil {
		res.Code = 0
	} else {
		res.Code = 1
	}
	res.Topic = get
	return
}

func (c *cTopic) Del(ctx context.Context, req *topic.DelReq) (res *topic.DelRes, err error) {
	err = service.Topic().Delete(ctx, req.Id)
	res = &topic.DelRes{}
	if err != nil {
		res.Id = 0
	}
	res.Id = req.Id
	return
}

func (c *cTopic) Edit(ctx context.Context, req *topic.EditReq) (res *topic.EditRes, err error) {
	id, err := service.Topic().Update(ctx, do.Topic{
		Id:               req.Id,
		PlatForm:         req.PlatForm,
		Topic:            req.Topic,
		FunctionDescribe: req.FunctionDescribe,
	})
	res = &topic.EditRes{Id: id}
	return
}

func (c *cTopic) SearchAllUpTopics(ctx context.Context, req *topic.SearchAllUpTopicReq) (res *topic.SearchAllUpTopicRes, err error) {
	get, err := service.Topic().GetAllUpTopics(ctx, req.UserId)
	res = &topic.SearchAllUpTopicRes{}
	if err != nil {
		res.Code = 0
	} else {
		res.Code = 1
	}
	res.Topic = get
	return
}

func (c *cTopic) SearchAllDownTopics(ctx context.Context, req *topic.SearchAllDownTopicReq) (res *topic.SearchAllDownTopicRes, err error) {
	get, err := service.Topic().GetAllDownTopics(ctx, req.UserId)
	res = &topic.SearchAllDownTopicRes{}
	if err != nil {
		res.Code = 0
	} else {
		res.Code = 1
	}
	res.Topic = get
	return
}

func (c *cTopic) SearchAllUpByDeviceIdTopics(ctx context.Context, req *topic.SearchAllUpTopicByDeviceIdReq) (res *topic.SearchAllUpTopicByDeviceIdRes, err error) {
	get, err := service.Topic().GetAllUpByDeviceIdTopics(ctx, req.UserId, req.DeviceId)
	res = &topic.SearchAllUpTopicByDeviceIdRes{}
	if err != nil {
		res.Code = 0
	} else {
		res.Code = 1
	}
	res.Topic = get
	return
}

func (c *cTopic) SearchAllDownByDeviceIdTopics(ctx context.Context, req *topic.SearchAllDownTopicByDeviceIdReq) (res *topic.SearchAllDownTopicByDeviceIdRes, err error) {
	get, err := service.Topic().GetAllDownByDeviceIdTopics(ctx, req.UserId, req.DeviceId)
	res = &topic.SearchAllDownTopicByDeviceIdRes{}
	if err != nil {
		res.Code = 0
	} else {
		res.Code = 1
	}
	res.Topic = get
	return
}
