package chatSupport

import (
	"cloudCustomerService/internal/consts"
	"cloudCustomerService/internal/middlewares"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"cloudCustomerService/api/chatSupport/message"
)

func (c *ControllerMessage) AcknowledgeMessages(ctx context.Context, req *message.AcknowledgeMessagesReq) (res *message.AcknowledgeMessagesRes, err error) {
	request := g.RequestFromCtx(ctx)
	userId := gconv.Int(middlewares.ChatSupportAuth().GetIdentity(request.Context()))
	ticket, err := service.Ticket().ChatSupportHasTicket(ctx, userId, req.TicketId)
	if ticket == nil || err != nil {
		return nil, consts.ChatSupportNoAuth
	}

	quantity, err := service.Message().UpdateMessageReadStatus(ctx, req.TicketId, consts.SendMessageToChatSupport, req.MessageIds)
	if err != nil {
		return nil, err
	}
	if quantity != 0 {
		_, err = service.Ticket().DecrementCsUnreadQuantity(ctx, req.TicketId, quantity)

	}
	return res, err
}
