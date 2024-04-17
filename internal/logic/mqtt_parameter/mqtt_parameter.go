package mqtt_parameter

import (
	"context"
	"iot_device_simulation/internal/dao"
	"iot_device_simulation/internal/model/do"
	"iot_device_simulation/internal/model/entity"
	"iot_device_simulation/internal/service"
)

func init() {
	service.RegisterMqtt(New())
}

func New() *iMqtt {
	return &iMqtt{}
}

type iMqtt struct {
}

func (i *iMqtt) Insert(ctx context.Context, deviceId int, parameter do.MqttParameter) (id int, err error) {
	result, err := dao.MqttParameter.Ctx(ctx).Data(&parameter).Insert()
	insertId, err := result.LastInsertId()
	id = int(insertId)
	_, err = dao.Device.Ctx(ctx).Where("id", deviceId).Data(do.Device{Id: deviceId, MqttParameterId: id}).OmitEmptyData().Update()
	return
}

func (i *iMqtt) Get(ctx context.Context, id int) (mqtt entity.MqttParameter, err error) {
	err = dao.MqttParameter.Ctx(ctx).Where(dao.MqttParameter.Columns().Id, id).Scan(&mqtt)
	return
}
func (i *iMqtt) Update(ctx context.Context, parameter do.MqttParameter) (id int, err error) {
	result, err := dao.MqttParameter.Ctx(ctx).Where(dao.MqttParameter.Columns().Id, parameter.Id).Data(parameter).OmitEmptyData().Update()
	insertId, err := result.LastInsertId()
	id = int(insertId)
	return
}
func (i *iMqtt) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.MqttParameter.Ctx(ctx).Where(dao.MqttParameter.Columns().Id, id).Delete()
	return
}
