// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SubTopicDao is the data access object for table sub_topic.
type SubTopicDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns SubTopicColumns // columns contains all the column names of Table for convenient usage.
}

// SubTopicColumns defines and stores column names for table sub_topic.
type SubTopicColumns struct {
	Id       string //
	SubTopic string // 订阅通信topic
	DeviceId string // 设备id
	State    string // 状态，0关闭，1开启
	WsParam  string // 开启ws的参数
	TopicId  string // 通信id
}

// subTopicColumns holds the columns for table sub_topic.
var subTopicColumns = SubTopicColumns{
	Id:       "id",
	SubTopic: "sub_topic",
	DeviceId: "device_id",
	State:    "state",
	WsParam:  "ws_param",
	TopicId:  "topic_id",
}

// NewSubTopicDao creates and returns a new DAO object for table data access.
func NewSubTopicDao() *SubTopicDao {
	return &SubTopicDao{
		group:   "default",
		table:   "sub_topic",
		columns: subTopicColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SubTopicDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SubTopicDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SubTopicDao) Columns() SubTopicColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SubTopicDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SubTopicDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SubTopicDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
