package transducers

import (
	"context"
	"encoding/json"
	"iot_device_simulation/api/transducers"
	"iot_device_simulation/internal/model/entity"
	"iot_device_simulation/internal/service"
)

var TransducersController = &cTransducers{}

type cTransducers struct {
}

func (c *cTransducers) Add(ctx context.Context, req *transducers.AddTransducersReq) (res *transducers.AddTransducersRes, err error) {
	marshal, err := json.Marshal(req.Options)
	if err != nil {
		return
	}
	err = service.Transducers().Insert(ctx,
		entity.Transducers{
			UserId:          req.UserId,
			Identifier:      req.Identifier,
			TransducersType: req.TransducersType,
			Option:          string(marshal),
		})
	res = &transducers.AddTransducersRes{}
	if err != nil {
		res.Code = 0
		return
	}
	res.Code = 1
	return
}

func (c *cTransducers) GetTransducersByDeviceId(ctx context.Context, req *transducers.GetTransducersByDeviceIdReq) (res *transducers.GetTransducersByDeviceIdRes, err error) {
	allTransducers, err := service.Transducers().GetAllTransducersByDeviceId(ctx, req.DeviceId)
	res = &transducers.GetTransducersByDeviceIdRes{Transducers: allTransducers}
	if err != nil {
		res.Code = 0
		return
	}
	res.Code = 1
	return
}

func (c *cTransducers) GetTransducersByUserIdReq(ctx context.Context, req *transducers.GetTransducersByUserIdReq) (res *transducers.GetTransducersByUserIdRes, err error) {
	page, err := service.Transducers().GetAllTransducersByUidPage(ctx, req.UserId, req.Page, req.Size)
	count, err := service.Transducers().GetAllTransducersCountByUid(ctx, req.UserId)
	res = &transducers.GetTransducersByUserIdRes{Transducers: page, Count: count}
	if err != nil {
		res.Code = 0
		return
	}
	res.Code = 1
	return
}

func (c *cTransducers) GetTransducersById(ctx context.Context, req *transducers.GetTransducerByIdReq) (res *transducers.GetTransducerByIdRes, err error) {
	transducer, err := service.Transducers().Get(ctx, req.Id)
	//transducer.Option = strings.Replace(transducer.Option, "\"", "'", -1)
	if err != nil {
		return
	}
	res = &transducers.GetTransducerByIdRes{Transducers: transducer}
	if err != nil {
		res.Code = 0
		return
	}
	res.Code = 1
	return
}

func (c *cTransducers) UpdateTransducersById(ctx context.Context, req *transducers.UpdateTransducersByDeviceIdReq) (res *transducers.UpdateTransducersByDeviceIdRes, err error) {
	marshal, err := json.Marshal(req.Options)
	if err != nil {
		return
	}
	err = service.Transducers().Update(ctx,
		entity.Transducers{
			Id:         req.Id,
			Identifier: req.Identifier,
			Option:     string(marshal),
		})
	res = &transducers.UpdateTransducersByDeviceIdRes{}
	if err != nil {
		res.Code = 0
		return
	}
	res.Code = 1
	return
}

func (c *cTransducers) DeleteTransducersById(ctx context.Context, req *transducers.DeleteTransducersByIdReq) (res *transducers.DeleteTransducersByIdRes, err error) {
	err = service.Transducers().Delete(ctx, req.Id)
	res = &transducers.DeleteTransducersByIdRes{}
	if err != nil {
		res.Code = 0
		return
	}
	res.Code = 1
	return
}

func (c *cTransducers) SetDeviceId(ctx context.Context, req *transducers.SetDeviceIdReq) (res *transducers.SetDeviceIdRes, err error) {
	err = service.Transducers().SetDeviceId(ctx, req.Id, req.DeviceId)
	res = &transducers.SetDeviceIdRes{}
	if err != nil {
		res.Code = 0
		return
	}
	res.Code = 1
	return
}

func (c *cTransducers) GetTransducersByUserId(ctx context.Context, req *transducers.GetTransducersByUidButNotDeviceReq) (res *transducers.GetTransducersByUidButNotDeviceRes, err error) {
	id, err := service.Transducers().GetAllTransducersByUidButDeviceId(ctx, req.UserId)
	res = &transducers.GetTransducersByUidButNotDeviceRes{Transducers: id}
	if err != nil {
		res.Code = 0
		return
	}
	res.Code = 1
	return
}
