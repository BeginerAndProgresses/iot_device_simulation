package transducers

import (
	"context"
	"encoding/json"
	"iot_device_simulation/internal/dao"
	"iot_device_simulation/internal/model/entity"
	"iot_device_simulation/internal/service"
	tu "iot_device_simulation/utility/transducers"
)

func init() {
	service.RegisterTransducers(New())
}

type iTransducers struct {
}

func New() *iTransducers {
	return &iTransducers{}
}

func (i *iTransducers) Insert(ctx context.Context, transducers entity.Transducers) error {
	_, err := dao.Transducers.Ctx(ctx).Insert(transducers)
	return err
}

// StartAllTransducers 开启所有传感器
func (i *iTransducers) StartAllTransducers(ctx context.Context, deviceId int, interval int64) error {
	transducers, err := i.GetAllTransducersByDeviceId(ctx, deviceId)
	if err != nil {
		return err
	}
	for _, transducer := range transducers {
		switch transducer.TransducersType {
		case tu.GasT:
			var intOption TransducerIntegerOption
			err = json.Unmarshal([]byte(transducer.Option), &intOption)
			if err != nil {
				continue
			}
			//	气体传感器
			gas := tu.NewGas(deviceId, intOption.Min, intOption.Max, transducer.Identifier, intOption.Unit)
			gas.Tr.SetInterval(interval)
			tu.CTransducers.Register(gas.Tr)
		case tu.HumidityT:
			var floatOption TransducerFloatOption
			err = json.Unmarshal([]byte(transducer.Option), &floatOption)
			if err != nil {
				continue
			}
			//	湿度传感器
			hum := tu.NewHumidity(floatOption.Min, floatOption.Max, transducer.Identifier, floatOption.Unit, deviceId)
			hum.Tr.SetInterval(interval)
			tu.CTransducers.Register(hum.Tr)
		case tu.LightT:
			var floatOption TransducerFloatOption
			err = json.Unmarshal([]byte(transducer.Option), &floatOption)
			if err != nil {
				continue
			}
			//	光敏传感器
			light := tu.NewLight(floatOption.Min, floatOption.Max, transducer.Identifier, floatOption.Unit, deviceId)
			light.Tr.SetInterval(interval)
			tu.CTransducers.Register(light.Tr)
		case tu.InfraredT:
			var boolOption TransducerBoolOption
			err = json.Unmarshal([]byte(transducer.Option), &boolOption)
			if err != nil {
				continue
			}
			//	红外传感器
			inf := tu.NewInfrared(transducer.Identifier, tu.BoolTypeStruct{Bool: 0, AbleDescribe: boolOption.AbleDescription, DisAbleDescribe: boolOption.DisAbleDescription}, deviceId)
			inf.Tr.SetInterval(interval)
			tu.CTransducers.Register(inf.Tr)
		case tu.TemperatureT:
			var floatOption TransducerFloatOption
			err = json.Unmarshal([]byte(transducer.Option), &floatOption)
			if err != nil {
				continue
			}
			//	温度传感器
			temp := tu.NewTemperature(floatOption.Min, floatOption.Max, transducer.Identifier, floatOption.Unit, deviceId)
			temp.Tr.SetInterval(interval)
			tu.CTransducers.Register(temp.Tr)
		case tu.PressureT:
			var floatOption TransducerFloatOption
			err = json.Unmarshal([]byte(transducer.Option), &floatOption)
			if err != nil {
				continue
			}
			//	压力传感器
			press := tu.NewPressure(floatOption.Min, floatOption.Max, transducer.Identifier, floatOption.Unit, deviceId)
			press.Tr.SetInterval(interval)
			tu.CTransducers.Register(press.Tr)
		case tu.SpeedT:
			var floatOption TransducerFloatOption
			err = json.Unmarshal([]byte(transducer.Option), &floatOption)
			if err != nil {
				continue
			}
			//	压力传感器
			speed := tu.NewSpeed(floatOption.Min, floatOption.Max, transducer.Identifier, floatOption.Unit, deviceId)
			speed.Tr.SetInterval(interval)
			tu.CTransducers.Register(speed.Tr)
		case tu.SmokeT:
			var enmuOption TransducerEnumOption
			err = json.Unmarshal([]byte(transducer.Option), &enmuOption)
			if err != nil {
				continue
			}
			//	烟雾传感器
			enmus := make([]interface{}, len(enmuOption.Enum))
			for k, _ := range enmuOption.Enum {
				enmus = append(enmus, k)
			}
			smoke := tu.NewSmoke(transducer.Identifier, tu.EnumTypeStruct{Enum: enmus, Describe: enmuOption.Enum}, deviceId)
			smoke.Tr.SetInterval(interval)
			tu.CTransducers.Register(smoke.Tr)
		case tu.LiquidLevelT:
			var floatOption TransducerFloatOption
			err = json.Unmarshal([]byte(transducer.Option), &floatOption)
			if err != nil {
				continue
			}
			//	温度传感器
			liquid := tu.NewLiquidLevel(floatOption.Min, floatOption.Max, transducer.Identifier, floatOption.Unit, deviceId)
			liquid.Tr.SetInterval(interval)
			tu.CTransducers.Register(liquid.Tr)
		}
	}
	return nil
}

func (i *iTransducers) StopAllTransducers(ctx context.Context, deviceId int) error {
	tu.CTransducers.RangeAll(func(t tu.Transducers, value bool) bool {
		if t.GetId() == deviceId {
			tu.CTransducers.UnRegister(t)
		}
		return true
	})
	return nil
}

func (i *iTransducers) GetAllTransducersByDeviceId(ctx context.Context, deviceId int) ([]entity.Transducers, error) {
	var transducers []entity.Transducers
	err := dao.Transducers.Ctx(ctx).Where(dao.Transducers.Columns().DeviceId, deviceId).Scan(&transducers)
	return transducers, err
}

func (i *iTransducers) GetAllTransducersByUidPage(ctx context.Context, uid int, page int, size int) (transducers []entity.Transducers, err error) {
	err = dao.Transducers.Ctx(ctx).Where(dao.Transducers.Columns().UserId, uid).Page(page, size).Scan(&transducers)
	return
}

func (i *iTransducers) GetAllTransducersCountByUid(ctx context.Context, uid int) (count int, err error) {
	count, err = dao.Transducers.Ctx(ctx).Where(dao.Transducers.Columns().UserId, uid).Count()
	return
}

func (i *iTransducers) Get(ctx context.Context, id int) (entity.Transducers, error) {
	var transducers entity.Transducers
	err := dao.Transducers.Ctx(ctx).Where(dao.Transducers.Columns().Id, id).Scan(&transducers)
	return transducers, err
}

func (i *iTransducers) Update(ctx context.Context, transducers entity.Transducers) error {
	_, err := dao.Transducers.Ctx(ctx).Where(dao.Transducers.Columns().Id, transducers.Id).OmitEmptyData().Update(transducers)
	return err
}

func (i *iTransducers) Delete(ctx context.Context, id int) error {
	_, err := dao.Transducers.Ctx(ctx).Where(dao.Transducers.Columns().Id, id).Delete()
	return err
}

func (i *iTransducers) SetDeviceId(ctx context.Context, id int, deviceId int) error {
	var transducers = entity.Transducers{
		Id:       id,
		DeviceId: uint(deviceId),
	}
	_, err := dao.Transducers.Ctx(ctx).Where(dao.Transducers.Columns().Id, id).OmitEmptyData().Update(&transducers)
	return err
}

func (i *iTransducers) GetAllTransducersByUidButDeviceId(ctx context.Context, uid int) ([]entity.Transducers, error) {
	var transducers []entity.Transducers
	err := dao.Transducers.Ctx(ctx).Where(dao.Transducers.Columns().UserId, uid).Where(dao.Transducers.Columns().DeviceId, 0).Scan(&transducers)
	return transducers, err
}
