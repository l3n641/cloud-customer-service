package chatSupport

import (
	"cloudCustomerService/api"
	"cloudCustomerService/internal/consts"
	"cloudCustomerService/internal/middlewares"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"cloudCustomerService/api/chatSupport/message"
)

func (c *ControllerMessage) GetMessageList(ctx context.Context, req *message.GetMessageListReq) (res *message.GetMessageListRes, err error) {
	request := g.RequestFromCtx(ctx)
	userId := gconv.Int(middlewares.ChatSupportAuth().GetIdentity(request.Context()))
	ticket, err := service.Ticket().ChatSupportHasTicket(ctx, userId, req.TicketId)
	if ticket == nil || err != nil {
		return nil, consts.ChatSupportNoAuth
	}

	out, err := service.Message().GetMessageList(ctx, &model.MessageListInput{
		Page:          req.Page,
		Size:          req.PageSize,
		LastMessageId: req.LastMessageId,
		TicketId:      ticket.Id,
	})

	service.Message().ReadMessage(ctx, ticket.Id, consts.SendMessageToChatSupport)
	service.Ticket().UpdateTicketByChatSupport(ctx, ticket.Id)

	return &message.GetMessageListRes{
		PaginationRes: api.PaginationRes{
			Item:     out.List,
			Total:    out.Total,
			Page:     out.Page,
			PageSize: out.Size,
		},
	}, err
}
