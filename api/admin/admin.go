// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package admin

import (
	"context"
	
	"cloudCustomerService/api/admin/account"
	"cloudCustomerService/api/admin/customerServiceAgent"
	"cloudCustomerService/api/admin/session"
)

type IAdminAccount interface {
	UserInfo(ctx context.Context, req *account.UserInfoReq) (res *account.UserInfoRes, err error)
}

type IAdminCustomerServiceAgent interface {
	CreateCustomerServiceAgents(ctx context.Context, req *customerServiceAgent.CreateCustomerServiceAgentsReq) (res *customerServiceAgent.CreateCustomerServiceAgentsRes, err error)
}

type IAdminSession interface {
	Login(ctx context.Context, req *session.LoginReq) (res *session.LoginRes, err error)
}


