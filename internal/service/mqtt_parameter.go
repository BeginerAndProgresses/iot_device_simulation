package service

import (
	"context"
	"fmt"
	"iot_device_simulation/internal/model/do"
	"iot_device_simulation/internal/model/entity"
)

type IMqtt interface {
	Insert(ctx context.Context, deviceId int, parameter do.MqttParameter) (id int, err error)
	Get(ctx context.Context, id int) (mqtt entity.MqttParameter, err error)
	Update(ctx context.Context, parameter do.MqttParameter) (id int, err error)
	Delete(ctx context.Context, id int) (err error)
	GetByDeviceId(ctx context.Context, deviceId int) (mqtt entity.MqttParameter, err error)
}

var localMqtt IMqtt

func Mqtt() IMqtt {
	if localMqtt == nil {
		fmt.Println("localMqtt未实现或是注册失败")
	}
	return localMqtt
}

func RegisterMqtt(i IMqtt) {
	localMqtt = i
}
