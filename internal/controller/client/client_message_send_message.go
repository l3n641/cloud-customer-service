package client

import (
	"cloudCustomerService/api/client/message"
	"cloudCustomerService/internal/consts"
	"cloudCustomerService/internal/middlewares"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerMessage) SendMessage(ctx context.Context, req *message.SendMessageReq) (res *message.SendMessageRes, err error) {
	request := g.RequestFromCtx(ctx)
	clientId := gconv.Int(middlewares.ClientAuth().GetIdentity(request.Context()))
	client, err := service.Client().GetClientById(ctx, clientId)
	if err != nil || client == nil {
		return nil, consts.ClientNotAuth
	}

	ticket, err := service.Ticket().HasTicket(ctx, clientId)
	if err != nil {
		return nil, consts.ClientNotTicket
	}

	var tickerId int
	var chatSupportId int

	if ticket == nil {
		chatSupport, err := service.ChatSupport().AssignChatSupport(ctx, client.MerchantId)

		if err != nil || chatSupport == nil {
			return nil, consts.NotHasChatSupport
		}
		chatSupportId = chatSupport.Id

		if tickerId, err = service.Ticket().CreateTicket(ctx, clientId, chatSupport.Id); err != nil {
			return nil, consts.ClientNotTicket
		}
	} else {
		tickerId = ticket.Id
		chatSupportId = ticket.ChatSupportId
	}

	service.Message().ClientSendMessage(ctx, &model.MessageAddInput{
		ClientId:      clientId,
		ChatSupportId: chatSupportId,
		Content:       req.Content,
		MsgType:       req.MsgType,
		SenderType:    consts.SendMessageToChatSupport,
		TicketId:      tickerId,
		RefererUrl:    req.RefererUrl,
	})

	service.Ticket().UpdateTicketByClient(ctx, tickerId)

	return &message.SendMessageRes{}, err

}
