package chatSupport

import (
	"cloudCustomerService/api"
	"cloudCustomerService/internal/middlewares"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"cloudCustomerService/api/chatSupport/ticket"
)

func (c *ControllerTicket) SearchTicket(ctx context.Context, req *ticket.SearchTicketReq) (res *ticket.SearchTicketRes, err error) {
	request := g.RequestFromCtx(ctx)
	userId := gconv.Int(middlewares.ChatSupportAuth().GetIdentity(request.Context()))
	out, err := service.Ticket().SearchTicket(ctx, &model.TicketSearchInput{
		ChatSupportId: userId,
		Page:          req.Page,
		Size:          req.PageSize,
	})

	return &ticket.SearchTicketRes{
		PaginationRes: api.PaginationRes{
			Item:     out.List,
			Total:    out.Total,
			Page:     out.Page,
			PageSize: out.Size,
		},
	}, err
}
