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
	ManageClientAccountBan(ctx context.Context, req *client.ManageClientAccountBanReq) (res *client.ManageClientAccountBanRes, err error)
}

type IChatSupportMessage interface {
	AcknowledgeMessages(ctx context.Context, req *message.AcknowledgeMessagesReq) (res *message.AcknowledgeMessagesRes, err error)
	GetMessageList(ctx context.Context, req *message.GetMessageListReq) (res *message.GetMessageListRes, err error)
	SubscriptionMsg(ctx context.Context, req *message.SubscriptionMsgReq) (res *message.SubscriptionMsgRes, err error)
	SendMessage(ctx context.Context, req *message.SendMessageReq) (res *message.SendMessageRes, err error)
}

type IChatSupportSession interface {
	Login(ctx context.Context, req *session.LoginReq) (res *session.LoginRes, err error)
	SignOut(ctx context.Context, req *session.SignOutReq) (res *session.SignOutRes, err error)
}

type IChatSupportTicket interface {
	DetailTicket(ctx context.Context, req *ticket.DetailTicketReq) (res *ticket.DetailTicketRes, err error)
	SearchTicket(ctx context.Context, req *ticket.SearchTicketReq) (res *ticket.SearchTicketRes, err error)
	UpdateTicketRemark(ctx context.Context, req *ticket.UpdateTicketRemarkReq) (res *ticket.UpdateTicketRemarkRes, err error)
}


