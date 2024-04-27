package device

import (
	"context"
	"fmt"
	"iot_device_simulation/internal/dao"
	"iot_device_simulation/internal/model/entity"
	"iot_device_simulation/internal/packed/conn"
	"iot_device_simulation/internal/service"
	tu "iot_device_simulation/utility/transducers"
	"sync"
	"time"
)

type iSimulateDevice struct {
}

func NewSimulateDevice() *iSimulateDevice {
	return &iSimulateDevice{}
}

func (i *iSimulateDevice) Insert(ctx context.Context, transducers entity.SimulateDevice) error {
	_, err := dao.SimulateDevice.Ctx(ctx).Insert(&transducers)
	if err != nil {
		return err
	}
	return nil
}

func (i *iSimulateDevice) Get(ctx context.Context, id int) (simulateDevice entity.SimulateDevice, err error) {
	err = dao.SimulateDevice.Ctx(ctx).Where(dao.SimulateDevice.Columns().Id, id).Scan(&simulateDevice)
	return
}

func (i *iSimulateDevice) Start(ctx context.Context, deviceId int) error {
	var simulate entity.SimulateDevice
	err2 := dao.SimulateDevice.Ctx(ctx).Where(dao.SimulateDevice.Columns().Id, deviceId).Scan(&simulate)
	if err2 != nil {
		return err2
	}
	var transducers []entity.Transducers
	// 找传感器
	err := dao.Transducers.Ctx(ctx).Where(dao.Transducers.Columns().DeviceId, deviceId).Scan(&transducers)
	if err != nil {
		return err
	}
	err = service.Transducers().StartAllTransducers(ctx, deviceId, int64(simulate.Intervals))
	if err != nil {
		return err
	}
	// 开启Mqtt连接
	conn.CreatAConn(ctx, deviceId)
	_, err = dao.SimulateDevice.Ctx(ctx).Where(dao.SimulateDevice.Columns().Id, deviceId).OmitEmptyData().Update(&entity.SimulateDevice{State: 2})
	if err != nil {
		return err
	}
	device, err := service.SimulateDevice().Get(ctx, deviceId)
	if err != nil {
		return err
	}
	getTopic := simulate.Topic
	if err != nil {
		return err
	}
	sendTopic := conn.TopicPadding(device.PlatForm, getTopic, device.DeviceName, device.ProductId, device.DeviceId)
	go func() {
		ticker := time.NewTicker(time.Second * time.Duration(simulate.Intervals))
		var infoMap = make(map[string]interface{}, 10)
		var infoMapLooker = sync.Mutex{}
		for {
			select {
			case info := <-tu.CTransducers.Infos:
				infoMapLooker.Lock()
				infoMap[info.Identify] = info.Data
				infoMapLooker.Unlock()
				fmt.Println("infoMap", infoMap)
			case <-ticker.C:
				// 发送数据
				fmt.Println("发送数据到平台")
				switch device.PlatForm {
				case "腾讯云":
					conn.CManager.PublishToTencent(deviceId, sendTopic, infoMap)
				case "华为云":
					conn.CManager.PublishToTencent(deviceId, sendTopic, infoMap)
				case "阿里云":
					conn.CManager.PublishToAli(deviceId, sendTopic, infoMap)
				}
			case connClose := <-conn.CManager.CloseChan:
				if connClose.DeviceId == device.Id {
					fmt.Println("连接停止")
					conn.CManager.DelConnect(connClose)
					return
				}
			}
		}
	}()
	return nil
}

func (i *iSimulateDevice) Stop(ctx context.Context, deviceId int) error {
	// 关闭所有的传感器
	err := service.Transducers().StopAllTransducers(ctx, deviceId)
	err = service.SimulateDevice().Update(ctx, entity.SimulateDevice{Id: deviceId, State: 1})
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	conn.CloseAConn(deviceId)
	return nil
}

func (i *iSimulateDevice) GetByUid(ctx context.Context, userid int) (simulateDevice []entity.SimulateDevice, err error) {
	err = dao.SimulateDevice.Ctx(ctx).Where(dao.SimulateDevice.Columns().UserId, userid).Scan(&simulateDevice)
	return
}

func (i *iSimulateDevice) Update(ctx context.Context, transducers entity.SimulateDevice) error {
	_, err := dao.SimulateDevice.Ctx(ctx).Where(dao.SimulateDevice.Columns().Id, transducers.Id).OmitEmptyData().Update(&transducers)
	if err != nil {
		return err
	}
	return nil
}

func (i *iSimulateDevice) Delete(ctx context.Context, id int) error {
	_, err := dao.SimulateDevice.Ctx(ctx).Where(dao.SimulateDevice.Columns().Id, id).Delete()
	if err != nil {
		return err
	}
	return nil
}
