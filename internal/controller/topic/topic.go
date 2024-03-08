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
		PlatForm:         req.PlatForm,
		Topic:            req.Topic,
		FunctionDescribe: req.FunctionDescribe,
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
