package client

import (
	"cloudCustomerService/api"
	"cloudCustomerService/internal/consts"
	"cloudCustomerService/internal/middlewares"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"cloudCustomerService/api/client/message"
)

func (c *ControllerMessage) GetMessageList(ctx context.Context, req *message.GetMessageListReq) (res *message.GetMessageListRes, err error) {
	request := g.RequestFromCtx(ctx)
	clientId := gconv.Int(middlewares.ClientAuth().GetIdentity(request.Context()))
	ticket, err := service.Ticket().HasTicket(ctx, clientId)
	if err != nil || ticket == nil {
		return &message.GetMessageListRes{}, err
	}
	out, err := service.Message().ClientGetMessageList(ctx, &model.ClientMessageListInput{
		Page:          req.Page,
		Size:          req.PageSize,
		LastMessageId: req.LastMessageId,
		TicketId:      ticket.Id,
		OnlyUnread:    req.OnlyUnread,
	})

	service.Message().ReadMessage(ctx, ticket.Id, consts.SendMessageToClient)

	return &message.GetMessageListRes{
		PaginationRes: api.PaginationRes{
			Item:     out.List,
			Total:    out.Total,
			Page:     out.Page,
			PageSize: out.Size,
		},
	}, err
}
