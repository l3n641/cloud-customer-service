// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package chatSupport

import (
	"context"
	
	"cloudCustomerService/api/chatSupport/account"
	"cloudCustomerService/api/chatSupport/session"
)

type IChatSupportAccount interface {
	UserInfo(ctx context.Context, req *account.UserInfoReq) (res *account.UserInfoRes, err error)
}

type IChatSupportSession interface {
	Login(ctx context.Context, req *session.LoginReq) (res *session.LoginRes, err error)
	SignOut(ctx context.Context, req *session.SignOutReq) (res *session.SignOutRes, err error)
}


