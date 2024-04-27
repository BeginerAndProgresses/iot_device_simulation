// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TransducersDao is the data access object for table transducers.
type TransducersDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns TransducersColumns // columns contains all the column names of Table for convenient usage.
}

// TransducersColumns defines and stores column names for table transducers.
type TransducersColumns struct {
	Id              string //
	UserId          string // 用户id
	DeviceId        string // 设备id
	TransducersType string // 传感器类型
	Identifier      string // 标识符
	Option          string // 配置，传入json
}

// transducersColumns holds the columns for table transducers.
var transducersColumns = TransducersColumns{
	Id:              "id",
	UserId:          "user_id",
	DeviceId:        "device_id",
	TransducersType: "transducers_type",
	Identifier:      "identifier",
	Option:          "option",
}

// NewTransducersDao creates and returns a new DAO object for table data access.
func NewTransducersDao() *TransducersDao {
	return &TransducersDao{
		group:   "default",
		table:   "transducers",
		columns: transducersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TransducersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TransducersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TransducersDao) Columns() TransducersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TransducersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TransducersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TransducersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
