// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MqttParameterDao is the data access object for table mqtt_parameter.
type MqttParameterDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns MqttParameterColumns // columns contains all the column names of Table for convenient usage.
}

// MqttParameterColumns defines and stores column names for table mqtt_parameter.
type MqttParameterColumns struct {
	Id            string //
	ClientId      string // Client ID
	Port          string // ç«¯å£å·
	ServerAddress string // æœåŠ¡å™¨åœ°å€
	Username      string // username
	Password      string // password
	UserId        string // ç”¨æˆ·id
}

// mqttParameterColumns holds the columns for table mqtt_parameter.
var mqttParameterColumns = MqttParameterColumns{
	Id:            "id",
	ClientId:      "client_id",
	Port:          "port",
	ServerAddress: "server_address",
	Username:      "username",
	Password:      "password",
	UserId:        "user_id",
}

// NewMqttParameterDao creates and returns a new DAO object for table data access.
func NewMqttParameterDao() *MqttParameterDao {
	return &MqttParameterDao{
		group:   "default",
		table:   "mqtt_parameter",
		columns: mqttParameterColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MqttParameterDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MqttParameterDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MqttParameterDao) Columns() MqttParameterColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MqttParameterDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MqttParameterDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MqttParameterDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
