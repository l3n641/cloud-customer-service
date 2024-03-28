package client

import (
	"cloudCustomerService/internal/consts"
	"cloudCustomerService/internal/middlewares"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"cloudCustomerService/api/client/message"
)

func (c *ControllerMessage) GetUnreadMessageQuantity(ctx context.Context, req *message.GetUnreadMessageQuantityReq) (res *message.GetUnreadMessageQuantityRes, err error) {
	request := g.RequestFromCtx(ctx)
	clientId := gconv.Int(middlewares.ClientAuth().GetIdentity(request.Context()))
	ticket, err := service.Ticket().HasTicket(ctx, clientId)
	if err != nil || ticket == nil {
		return &message.GetUnreadMessageQuantityRes{}, err
	}

	quantity, err := service.Message().GetUnreadMessageQuantity(ctx, consts.SendMessageToClient, ticket.Id)
	return &message.GetUnreadMessageQuantityRes{
		Quantity: quantity,
	}, err

}
