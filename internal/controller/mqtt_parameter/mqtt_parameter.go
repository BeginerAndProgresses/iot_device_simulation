package mqtt_parameter

import (
	"context"
	"fmt"
	"iot_device_simulation/api/mqtt_parameter"
	"iot_device_simulation/internal/model/do"
	"iot_device_simulation/internal/packed/util"
	"iot_device_simulation/internal/service"
)

var MqttController = &cMqtt{}

type cMqtt struct {
}

func (c *cMqtt) Add(ctx context.Context, req *mqtt_parameter.AddReq) (res *mqtt_parameter.AddRes, err error) {
	fmt.Println("Addmqtt: ", req)
	id, err := service.Mqtt().Insert(ctx, req.DeviceId, do.MqttParameter{
		ClientId:      req.ClientId,
		Port:          req.Port,
		ServerAddress: req.ServerAddress,
		Username:      req.Username,
		Password:      req.Password,
		UserId:        req.UserId,
	})
	res = &mqtt_parameter.AddRes{Id: id}
	return
}

func (c *cMqtt) Edit(ctx context.Context, req *mqtt_parameter.EditReq) (res *mqtt_parameter.EditRes, err error) {
	id, err := service.Mqtt().Update(ctx, do.MqttParameter{
		Id:            req.Id,
		ClientId:      req.ClientId,
		Port:          req.Port,
		ServerAddress: req.ServerAddress,
		Username:      req.Username,
		Password:      req.Password,
	})
	res = &mqtt_parameter.EditRes{Id: id}
	return
}

// Del 失败返回0
func (c *cMqtt) Del(ctx context.Context, req *mqtt_parameter.DelReq) (res *mqtt_parameter.DelRes, err error) {
	fmt.Println("Delmqtt: ", req)
	err = service.Mqtt().Delete(ctx, req.Id)
	res = &mqtt_parameter.DelRes{}
	if err != nil {
		res.Id = 0
	}
	res.Id = req.Id
	return
}

func (c *cMqtt) Search(ctx context.Context, req *mqtt_parameter.SearchReq) (res *mqtt_parameter.SearchRes, err error) {
	fmt.Println("Searchmqtt: ", req)
	mqtt, err := service.Mqtt().Get(ctx, req.Id)
	res = &mqtt_parameter.SearchRes{}
	if err != nil {
		res.Code = 0
	} else {
		res.Code = 1
	}
	err = util.Transfer(&mqtt, &res.Mqtt)
	return
}

func (c *cMqtt) SearchByDeviceId(ctx context.Context, req *mqtt_parameter.SearchByDeviceIdReq) (res *mqtt_parameter.SearchByDeviceIdRes, err error) {
	mqtt, err := service.Mqtt().GetByDeviceId(ctx, req.DeviceId)
	res = &mqtt_parameter.SearchByDeviceIdRes{}
	if err != nil {
		res.Code = 0
		return
	}
	res.Code = 1
	res.Mqtt = mqtt
	return
}
