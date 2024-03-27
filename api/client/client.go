// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package client

import (
	"context"
	
	"cloudCustomerService/api/client/session"
)

type IClientSession interface {
	Login(ctx context.Context, req *session.LoginReq) (res *session.LoginRes, err error)
}


