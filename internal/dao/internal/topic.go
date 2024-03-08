// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TopicDao is the data access object for table topic.
type TopicDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns TopicColumns // columns contains all the column names of Table for convenient usage.
}

// TopicColumns defines and stores column names for table topic.
type TopicColumns struct {
	Id               string // id
	PlatForm         string // 平台名
	Topic            string // 通信topic
	FunctionDescribe string // 功能描述
}

// topicColumns holds the columns for table topic.
var topicColumns = TopicColumns{
	Id:               "id",
	PlatForm:         "plat_form",
	Topic:            "topic",
	FunctionDescribe: "function_describe",
}

// NewTopicDao creates and returns a new DAO object for table data access.
func NewTopicDao() *TopicDao {
	return &TopicDao{
		group:   "default",
		table:   "topic",
		columns: topicColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TopicDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TopicDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TopicDao) Columns() TopicColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TopicDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TopicDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TopicDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
