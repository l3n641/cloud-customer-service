package chatSupport

import (
	"cloudCustomerService/internal/consts"
	"cloudCustomerService/internal/middlewares"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"cloudCustomerService/api/chatSupport/message"
)

func (c *ControllerMessage) SendMessage(ctx context.Context, req *message.SendMessageReq) (res *message.SendMessageRes, err error) {
	request := g.RequestFromCtx(ctx)
	userId := gconv.Int(middlewares.ChatSupportAuth().GetIdentity(request.Context()))
	ticket, err := service.Ticket().ChatSupportHasTicket(ctx, userId, req.TicketId)
	if ticket == nil || err != nil {
		return nil, consts.ChatSupportNoAuth
	}

	msg, err := service.Message().SendMessage(ctx, &model.MessageAddInput{
		ClientId:      ticket.ClientId,
		ChatSupportId: userId,
		Content:       req.Content,
		MsgType:       req.MsgType,
		SenderType:    consts.SendMessageToClient,
		TicketId:      req.TicketId,
	})

	service.Ticket().UpdateTicketLastMessageTimeByChatSupport(ctx, req.TicketId)
	return &message.SendMessageRes{
		Id:         msg.Id,
		CreateAt:   msg.CreateAt,
		TicketId:   msg.TicketId,
		ClientId:   msg.ClientId,
		Content:    msg.Content,
		MsgType:    msg.MsgType,
		SenderType: msg.SenderType,
	}, err
}
