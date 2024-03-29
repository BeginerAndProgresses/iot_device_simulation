// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"iot_device_simulation/internal/dao/internal"
)

// internalMqttParameterDao is internal type for wrapping internal DAO implements.
type internalMqttParameterDao = *internal.MqttParameterDao

// mqttParameterDao is the data access object for table mqtt_parameter.
// You can define custom methods on it to extend its functionality as you wish.
type mqttParameterDao struct {
	internalMqttParameterDao
}

var (
	// MqttParameter is globally public accessible object for table mqtt_parameter operations.
	MqttParameter = mqttParameterDao{
		internal.NewMqttParameterDao(),
	}
)

// Fill with you ideas below.
