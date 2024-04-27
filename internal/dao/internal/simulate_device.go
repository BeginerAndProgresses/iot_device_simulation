// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SimulateDeviceDao is the data access object for table simulate_device.
type SimulateDeviceDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SimulateDeviceColumns // columns contains all the column names of Table for convenient usage.
}

// SimulateDeviceColumns defines and stores column names for table simulate_device.
type SimulateDeviceColumns struct {
	Id         string //
	PlatForm   string // 平台
	DeviceName string // 设备名
	DeviceId   string // 设备id
	State      string // 状态
	ProductId  string // 产品id
	UserId     string // 用户id
	Intervals  string // 时间间隔，默认20秒
	MqttOption string // mqtt连接参数
	Topic      string // 通信主题
}

// simulateDeviceColumns holds the columns for table simulate_device.
var simulateDeviceColumns = SimulateDeviceColumns{
	Id:         "id",
	PlatForm:   "plat_form",
	DeviceName: "device_name",
	DeviceId:   "device_id",
	State:      "state",
	ProductId:  "product_id",
	UserId:     "user_id",
	Intervals:  "intervals",
	MqttOption: "mqtt_option",
	Topic:      "topic",
}

// NewSimulateDeviceDao creates and returns a new DAO object for table data access.
func NewSimulateDeviceDao() *SimulateDeviceDao {
	return &SimulateDeviceDao{
		group:   "default",
		table:   "simulate_device",
		columns: simulateDeviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SimulateDeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SimulateDeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SimulateDeviceDao) Columns() SimulateDeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SimulateDeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SimulateDeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SimulateDeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
