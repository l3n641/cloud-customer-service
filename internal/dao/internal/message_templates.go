// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MessageTemplatesDao is the data access object for table message_templates.
type MessageTemplatesDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns MessageTemplatesColumns // columns contains all the column names of Table for convenient usage.
}

// MessageTemplatesColumns defines and stores column names for table message_templates.
type MessageTemplatesColumns struct {
	Id       string //
	CreateAt string // 消息时间
	MsgType  string // 消息类型 1-欢迎用语,2-离线用语,3-在线用语
	Lang     string // 语种
	Content  string // 消息内容
}

// messageTemplatesColumns holds the columns for table message_templates.
var messageTemplatesColumns = MessageTemplatesColumns{
	Id:       "id",
	CreateAt: "create_at",
	MsgType:  "msg_type",
	Lang:     "lang",
	Content:  "content",
}

// NewMessageTemplatesDao creates and returns a new DAO object for table data access.
func NewMessageTemplatesDao() *MessageTemplatesDao {
	return &MessageTemplatesDao{
		group:   "default",
		table:   "message_templates",
		columns: messageTemplatesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MessageTemplatesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MessageTemplatesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MessageTemplatesDao) Columns() MessageTemplatesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MessageTemplatesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MessageTemplatesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MessageTemplatesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
