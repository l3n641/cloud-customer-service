// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MessagesDao is the data access object for table messages.
type MessagesDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns MessagesColumns // columns contains all the column names of Table for convenient usage.
}

// MessagesColumns defines and stores column names for table messages.
type MessagesColumns struct {
	Id            string //
	CreateAt      string // 消息时间
	TicketId      string // 工单id
	SenderType    string // 发送类型:1-客人发送给客服，2-客服发送给客人
	ClientId      string // 顾客的id
	ChatSupportId string // 客服的id
	Content       string // 消息内容
	IsRead        string // 消息是否被对方读取:0-未读，1-已读
	ReadTime      string // 读取消息的时间
}

// messagesColumns holds the columns for table messages.
var messagesColumns = MessagesColumns{
	Id:            "id",
	CreateAt:      "create_at",
	TicketId:      "ticket_id",
	SenderType:    "sender_type",
	ClientId:      "client_id",
	ChatSupportId: "chat_support_id",
	Content:       "content",
	IsRead:        "is_read",
	ReadTime:      "read_time",
}

// NewMessagesDao creates and returns a new DAO object for table data access.
func NewMessagesDao() *MessagesDao {
	return &MessagesDao{
		group:   "default",
		table:   "messages",
		columns: messagesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MessagesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MessagesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MessagesDao) Columns() MessagesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MessagesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MessagesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MessagesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
