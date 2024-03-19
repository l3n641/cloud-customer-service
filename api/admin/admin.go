// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package admin

import (
	"context"
	
	"cloudCustomerService/api/admin/session"
)

type IAdminSession interface {
	Login(ctx context.Context, req *session.LoginReq) (res *session.LoginRes, err error)
}


