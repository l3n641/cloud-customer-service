// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package chatSupport

import (
	"context"
	
	"cloudCustomerService/api/chatSupport/session"
)

type IChatSupportSession interface {
	Login(ctx context.Context, req *session.LoginReq) (res *session.LoginRes, err error)
}


