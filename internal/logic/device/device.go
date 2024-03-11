package device

// 用于接口代码的实现
import (
	"context"
	"fmt"
	"iot_device_simulation/internal/dao"
	"iot_device_simulation/internal/model/do"
	"iot_device_simulation/internal/model/entity"
	"iot_device_simulation/internal/packed/conn"
	"iot_device_simulation/internal/packed/util"
	"iot_device_simulation/internal/service"
)

func init() {
	service.RegisterDevice(New())
	// 将所用设备状态调整为未连接
}

func New() *iDevice {
	return &iDevice{}
}

type iDevice struct{}

// Get 获取单个设备
func (i *iDevice) Get(ctx context.Context, id int) (device entity.Device, err error) {
	err = dao.Device.Ctx(ctx).Where("id", id).Scan(&device)
	return
}

func (i *iDevice) GetByPlatform(ctx context.Context, platform string) (devices []entity.Device, err error) {
	err = dao.Device.Ctx(ctx).Where("plat_form", platform).Scan(&devices)
	return
}

// Insert 插入一个设备
func (i *iDevice) Insert(ctx context.Context, device do.Device) (id int, err error) {
	result, err := dao.Device.Ctx(ctx).Data(device).Insert()
	if err != nil {
		panic("mysql add error：" + err.Error())
	}
	insertId, err := result.LastInsertId()
	id = int(insertId)
	return
}

// Update 修改一个设备信息
func (i *iDevice) Update(ctx context.Context, device do.Device) (id int, err error) {
	result, err := dao.Device.Ctx(ctx).Where(dao.Device.Columns().Id, device.Id).Data(device).OmitEmptyData().Update()
	if err != nil {
		panic("mysql update error：" + err.Error())
	}
	insertId, err := result.LastInsertId()
	id = int(insertId)
	return
}

// Delete 删除一个设备
func (i *iDevice) Delete(ctx context.Context, id int) (err error) {
	_, err = dao.Device.Ctx(ctx).Where(dao.Device.Columns().Id, id).Delete()
	if err != nil {
		panic("mysql delete error：" + err.Error())
	}
	return
}

func (i *iDevice) ConnMqtt(ctx context.Context, deviceId int) (id, state int, err error) {
	//获取连接参数
	var (
		t_dev  entity.Device
		t_mqtt entity.MqttParameter
	)

	err = dao.Device.Ctx(ctx).Where("id", deviceId).Scan(&t_dev)
	if err != nil || t_dev.MqttParameterId == 0 {
		id = 0
		state = 0
		return
	}

	err = dao.MqttParameter.Ctx(ctx).Where("id", t_dev.MqttParameterId).Scan(&t_mqtt)
	if err != nil {
		id = 0
		state = 3
		return
	}
	err = conn.Conn(deviceId, &t_mqtt)
	if err != nil {
		id = 0
		state = 0
		return
	}
	id = deviceId
	state = 2
	_, err = dao.Device.Ctx(ctx).Where("id", deviceId).Data(do.Device{Id: deviceId, State: state}).OmitEmptyData().Update()
	if err != nil {
		fmt.Println("error updating device", err)
	}
	return
}

func (i *iDevice) DisConnMqtt(ctx context.Context, deviceId int) (id, state int, err error) {
	fmt.Println("调用DisConnMqtt")
	_, err = dao.Device.Ctx(ctx).Where("id", deviceId).Data(do.Device{Id: deviceId, State: 1}).OmitEmptyData().Update()
	if err != nil {
		id = 0
		fmt.Println("error updating device", err)
		return
	}
	id = deviceId
	err = conn.DisConn(deviceId)
	if err != nil {
		id = 0
		state = 0
		return
	}
	state = 1
	//_, err = dao.Device.Ctx(ctx).Where("id", deviceId).Data(&do.Device{Id: deviceId, State: 0}).OmitEmptyData().Update()
	if err != nil {
		id = 0
		fmt.Println("error updating device", err)
		return
	}
	return
}

func (i *iDevice) InfoPost(ctx context.Context, deviceId, topicId int, json string) (err error) {
	//上报数据
	//获取连接参数
	var (
		t_dev   entity.Device
		t_topic entity.Topic
		t_mqtt  entity.MqttParameter
		topic   string
	)
	err = dao.Device.Ctx(ctx).Where("id", deviceId).Scan(&t_dev)
	err = dao.Topic.Ctx(ctx).Where("id", topicId).Scan(&t_topic)
	m := make(map[string]string)
	switch t_topic.PlatForm {
	case "阿里云":
		m["deviceName"] = t_dev.DeviceName
		m["productKey"] = t_dev.ProductId
		topic = util.VariableString2String(t_topic.Topic, m, "${", "}")
	case "腾讯云":
		m["DeviceName"] = t_dev.DeviceName
		m["ProductID"] = t_dev.ProductId
		topic = util.VariableString2String(t_topic.Topic, m, "{", "}")
	case "OneNet":
	case "华为云":
		err = dao.MqttParameter.Ctx(ctx).Where("id", t_dev.MqttParameterId).Scan(&t_mqtt)
		m["device_id"] = t_mqtt.Username
		topic = util.VariableString2String(t_topic.Topic, m, "{", "}")
	}

	_, err = dao.PublishInfo.Ctx(ctx).Data(do.PublishInfo{Topic: topic, Json: json}).Insert()
	//id, err := result.LastInsertId()
	fmt.Printf("topic", topic)
	err = conn.Publish(deviceId, topic, json)
	return
}

// TopicSub 订阅topic
func (i *iDevice) TopicSub(ctx context.Context, deviceId, topicId int) (err error) {
	//上报数据
	//获取连接参数
	var (
		t_dev   entity.Device
		t_topic entity.Topic
		t_mqtt  entity.MqttParameter
		topic   string
	)
	err = dao.Device.Ctx(ctx).Where("id", deviceId).Scan(&t_dev)
	err = dao.Topic.Ctx(ctx).Where("id", topicId).Scan(&t_topic)
	m := make(map[string]string)
	switch t_topic.PlatForm {
	case "阿里云":
		m["deviceName"] = t_dev.DeviceName
		m["productKey"] = t_dev.ProductId
		topic = util.VariableString2String(t_topic.Topic, m, "${", "}")
	case "腾讯云":
		m["DeviceName"] = t_dev.DeviceName
		m["ProductID"] = t_dev.ProductId
		topic = util.VariableString2String(t_topic.Topic, m, "{", "}")
	case "OneNet":
	case "华为云":
		err = dao.MqttParameter.Ctx(ctx).Where("id", t_dev.MqttParameterId).Scan(&t_mqtt)
		m["device_id"] = t_mqtt.Username
		topic = util.VariableString2String(t_topic.Topic, m, "{", "}")
	}

	//_, err = dao.PublishInfo.Ctx(ctx).Data(do.PublishInfo{Topic: topic}).Insert()
	//id, err := result.LastInsertId()
	fmt.Printf("topic", topic)
	err = conn.Subscribe(deviceId, topic)
	return
}
