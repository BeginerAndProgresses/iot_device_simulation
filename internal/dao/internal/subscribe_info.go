// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SubscribeInfoDao is the data access object for table subscribe_info.
type SubscribeInfoDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SubscribeInfoColumns // columns contains all the column names of Table for convenient usage.
}

// SubscribeInfoColumns defines and stores column names for table subscribe_info.
type SubscribeInfoColumns struct {
	Id       string //
	SubName  string //
	Topic    string //
	Info     string // 返回的信息
	DeviceId string //
}

// subscribeInfoColumns holds the columns for table subscribe_info.
var subscribeInfoColumns = SubscribeInfoColumns{
	Id:       "id",
	SubName:  "sub_name",
	Topic:    "topic",
	Info:     "info",
	DeviceId: "device_id",
}

// NewSubscribeInfoDao creates and returns a new DAO object for table data access.
func NewSubscribeInfoDao() *SubscribeInfoDao {
	return &SubscribeInfoDao{
		group:   "default",
		table:   "subscribe_info",
		columns: subscribeInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SubscribeInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SubscribeInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SubscribeInfoDao) Columns() SubscribeInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SubscribeInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SubscribeInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SubscribeInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
