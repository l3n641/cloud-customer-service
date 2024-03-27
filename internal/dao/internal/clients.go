// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ClientsDao is the data access object for table clients.
type ClientsDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ClientsColumns // columns contains all the column names of Table for convenient usage.
}

// ClientsColumns defines and stores column names for table clients.
type ClientsColumns struct {
	Id                  string //
	CreateAt            string //
	Email               string //
	DialCode            string //
	Phone               string //
	MerchantId          string //
	Domain              string //
	LastLoginTime       string //
	LastSendMessageTime string //
	Ip                  string //
	BrowserFingerprint  string //
	Area                string //
	UserAgent           string //
	Lang                string //
	Status              string //
	IsOnline            string //
	Iso2                string //
	StatusChangeTime    string //
}

// clientsColumns holds the columns for table clients.
var clientsColumns = ClientsColumns{
	Id:                  "id",
	CreateAt:            "create_at",
	Email:               "email",
	DialCode:            "dial_code",
	Phone:               "phone",
	MerchantId:          "merchant_id",
	Domain:              "domain",
	LastLoginTime:       "last_login_time",
	LastSendMessageTime: "last_send_message_time",
	Ip:                  "ip",
	BrowserFingerprint:  "browser_fingerprint",
	Area:                "area",
	UserAgent:           "user_agent",
	Lang:                "lang",
	Status:              "status",
	IsOnline:            "is_online",
	Iso2:                "iso2",
	StatusChangeTime:    "status_change_time",
}

// NewClientsDao creates and returns a new DAO object for table data access.
func NewClientsDao() *ClientsDao {
	return &ClientsDao{
		group:   "default",
		table:   "clients",
		columns: clientsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ClientsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ClientsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ClientsDao) Columns() ClientsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ClientsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ClientsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ClientsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
