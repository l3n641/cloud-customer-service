// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MerchantsDao is the data access object for table merchants.
type MerchantsDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns MerchantsColumns // columns contains all the column names of Table for convenient usage.
}

// MerchantsColumns defines and stores column names for table merchants.
type MerchantsColumns struct {
	Id           string //
	MerchantId   string //
	MerchantName string //
}

// merchantsColumns holds the columns for table merchants.
var merchantsColumns = MerchantsColumns{
	Id:           "id",
	MerchantId:   "merchant_id",
	MerchantName: "merchant_name",
}

// NewMerchantsDao creates and returns a new DAO object for table data access.
func NewMerchantsDao() *MerchantsDao {
	return &MerchantsDao{
		group:   "default",
		table:   "merchants",
		columns: merchantsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MerchantsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MerchantsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MerchantsDao) Columns() MerchantsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MerchantsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MerchantsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MerchantsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
