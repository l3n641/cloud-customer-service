// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package client

import (
	"context"
	
	"cloudCustomerService/api/client/message"
	"cloudCustomerService/api/client/session"
)

type IClientMessage interface {
	GetUnreadMessageQuantity(ctx context.Context, req *message.GetUnreadMessageQuantityReq) (res *message.GetUnreadMessageQuantityRes, err error)
	SendMessage(ctx context.Context, req *message.SendMessageReq) (res *message.SendMessageRes, err error)
}

type IClientSession interface {
	Login(ctx context.Context, req *session.LoginReq) (res *session.LoginRes, err error)
}


