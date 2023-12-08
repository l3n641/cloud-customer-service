// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CustomerServiceAgentsDao is the data access object for table customer_service_agents.
type CustomerServiceAgentsDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns CustomerServiceAgentsColumns // columns contains all the column names of Table for convenient usage.
}

// CustomerServiceAgentsColumns defines and stores column names for table customer_service_agents.
type CustomerServiceAgentsColumns struct {
	Id          string //
	CreateAt    string // 创建时间
	Email       string // 邮箱
	Password    string // 加密后的密码
	Nickname    string // 昵称
	Avatar      string // 头像地址
	LastLoginAt string // 最后登录时间
	LastLoginIp string // 最后登录ip
	IsOnline    string // 是否在线 1-在线 0-不在线
}

// customerServiceAgentsColumns holds the columns for table customer_service_agents.
var customerServiceAgentsColumns = CustomerServiceAgentsColumns{
	Id:          "id",
	CreateAt:    "create_at",
	Email:       "email",
	Password:    "password",
	Nickname:    "nickname",
	Avatar:      "avatar",
	LastLoginAt: "last_login_at",
	LastLoginIp: "last_login_ip",
	IsOnline:    "is_online",
}

// NewCustomerServiceAgentsDao creates and returns a new DAO object for table data access.
func NewCustomerServiceAgentsDao() *CustomerServiceAgentsDao {
	return &CustomerServiceAgentsDao{
		group:   "default",
		table:   "customer_service_agents",
		columns: customerServiceAgentsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CustomerServiceAgentsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CustomerServiceAgentsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CustomerServiceAgentsDao) Columns() CustomerServiceAgentsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CustomerServiceAgentsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CustomerServiceAgentsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CustomerServiceAgentsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
