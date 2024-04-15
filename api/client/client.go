// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package client

import (
	"context"
	
	"cloudCustomerService/api/client/chatSupport"
	"cloudCustomerService/api/client/message"
	"cloudCustomerService/api/client/session"
)

type IClientChatSupport interface {
	IsOnlineChatSupport(ctx context.Context, req *chatSupport.IsOnlineChatSupportReq) (res *chatSupport.IsOnlineChatSupportRes, err error)
}

type IClientMessage interface {
	GetMessageList(ctx context.Context, req *message.GetMessageListReq) (res *message.GetMessageListRes, err error)
	GetUnreadMessageQuantity(ctx context.Context, req *message.GetUnreadMessageQuantityReq) (res *message.GetUnreadMessageQuantityRes, err error)
	SendMessage(ctx context.Context, req *message.SendMessageReq) (res *message.SendMessageRes, err error)
}

type IClientSession interface {
	Login(ctx context.Context, req *session.LoginReq) (res *session.LoginRes, err error)
}


