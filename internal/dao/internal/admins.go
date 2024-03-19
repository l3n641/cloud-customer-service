// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminsDao is the data access object for table admins.
type AdminsDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns AdminsColumns // columns contains all the column names of Table for convenient usage.
}

// AdminsColumns defines and stores column names for table admins.
type AdminsColumns struct {
	Id          string //
	CreateAt    string // 创建时间
	Username    string // 用户名
	Password    string // 加密后的密码
	Nickname    string // 昵称
	Avatar      string // 头像地址
	LastLoginAt string // 最后登录时间
	LastLoginIp string // 最后登录ip
}

// adminsColumns holds the columns for table admins.
var adminsColumns = AdminsColumns{
	Id:          "id",
	CreateAt:    "create_at",
	Username:    "username",
	Password:    "password",
	Nickname:    "nickname",
	Avatar:      "avatar",
	LastLoginAt: "last_login_at",
	LastLoginIp: "last_login_ip",
}

// NewAdminsDao creates and returns a new DAO object for table data access.
func NewAdminsDao() *AdminsDao {
	return &AdminsDao{
		group:   "default",
		table:   "admins",
		columns: adminsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminsDao) Columns() AdminsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
