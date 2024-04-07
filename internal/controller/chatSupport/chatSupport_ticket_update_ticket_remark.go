package chatSupport

import (
	"cloudCustomerService/internal/consts"
	"cloudCustomerService/internal/middlewares"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"cloudCustomerService/api/chatSupport/ticket"
)

func (c *ControllerTicket) UpdateTicketRemark(ctx context.Context, req *ticket.UpdateTicketRemarkReq) (res *ticket.UpdateTicketRemarkRes, err error) {
	request := g.RequestFromCtx(ctx)
	userId := gconv.Int(middlewares.ChatSupportAuth().GetIdentity(request.Context()))
	ticketData, err := service.Ticket().ChatSupportHasTicket(ctx, userId, req.Id)
	if ticketData == nil || err != nil {
		return nil, consts.ChatSupportNoAuth
	}
	_, err = service.Ticket().UpdateRemark(ctx, req.Id, req.Remark)
	return res, err

}
