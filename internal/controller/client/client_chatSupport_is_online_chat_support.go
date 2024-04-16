package client

import (
	"cloudCustomerService/internal/middlewares"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"cloudCustomerService/api/client/chatSupport"
)

func (c *ControllerChatSupport) IsOnlineChatSupport(ctx context.Context, req *chatSupport.IsOnlineChatSupportReq) (res *chatSupport.IsOnlineChatSupportRes, err error) {
	request := g.RequestFromCtx(ctx)
	clientId := gconv.Int(middlewares.ClientAuth().GetIdentity(request.Context()))
	ticket, err := service.Ticket().HasTicket(ctx, clientId)
	if err != nil || ticket == nil {
		return &chatSupport.IsOnlineChatSupportRes{IsOnline: -1, Messages: make([]model.MessageTemplates, 0)}, err
	}
	data, err := service.ChatSupport().GetOnlineStatus(ctx, ticket.ChatSupportId)

	return &chatSupport.IsOnlineChatSupportRes{IsOnline: data.IsOnline, Messages: data.Messages}, err
}
