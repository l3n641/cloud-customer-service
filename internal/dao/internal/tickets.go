// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TicketsDao is the data access object for table tickets.
type TicketsDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns TicketsColumns // columns contains all the column names of Table for convenient usage.
}

// TicketsColumns defines and stores column names for table tickets.
type TicketsColumns struct {
	Id               string //
	CreateAt         string //
	ClientId         string // 客人id
	ChatSupportId    string // 客服id
	LastMessageTime  string // 最后发送消息时间
	CsUnreadMsgCount string // 客服未读消息数
	Remark           string // 备注
}

// ticketsColumns holds the columns for table tickets.
var ticketsColumns = TicketsColumns{
	Id:               "id",
	CreateAt:         "create_at",
	ClientId:         "client_id",
	ChatSupportId:    "chat_support_id",
	LastMessageTime:  "last_message_time",
	CsUnreadMsgCount: "cs_unread_msg_count",
	Remark:           "remark",
}

// NewTicketsDao creates and returns a new DAO object for table data access.
func NewTicketsDao() *TicketsDao {
	return &TicketsDao{
		group:   "default",
		table:   "tickets",
		columns: ticketsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TicketsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TicketsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TicketsDao) Columns() TicketsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TicketsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TicketsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TicketsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
