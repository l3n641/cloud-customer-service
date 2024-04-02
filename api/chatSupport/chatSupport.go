// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package chatSupport

import (
	"context"
	
	"cloudCustomerService/api/chatSupport/account"
	"cloudCustomerService/api/chatSupport/client"
	"cloudCustomerService/api/chatSupport/message"
	"cloudCustomerService/api/chatSupport/session"
	"cloudCustomerService/api/chatSupport/ticket"
)

type IChatSupportAccount interface {
	UserInfo(ctx context.Context, req *account.UserInfoReq) (res *account.UserInfoRes, err error)
}

type IChatSupportClient interface {
	GetClientDetail(ctx context.Context, req *client.GetClientDetailReq) (res *client.GetClientDetailRes, err error)
}

type IChatSupportMessage interface {
	GetMessageList(ctx context.Context, req *message.GetMessageListReq) (res *message.GetMessageListRes, err error)
	SendMessage(ctx context.Context, req *message.SendMessageReq) (res *message.SendMessageRes, err error)
}

type IChatSupportSession interface {
	Login(ctx context.Context, req *session.LoginReq) (res *session.LoginRes, err error)
	SignOut(ctx context.Context, req *session.SignOutReq) (res *session.SignOutRes, err error)
}

type IChatSupportTicket interface {
	SearchTicket(ctx context.Context, req *ticket.SearchTicketReq) (res *ticket.SearchTicketRes, err error)
}


