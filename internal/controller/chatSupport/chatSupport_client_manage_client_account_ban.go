package chatSupport

import (
	"cloudCustomerService/internal/consts"
	"cloudCustomerService/internal/middlewares"
	"cloudCustomerService/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"cloudCustomerService/api/chatSupport/client"
)

func (c *ControllerClient) ManageClientAccountBan(ctx context.Context, req *client.ManageClientAccountBanReq) (res *client.ManageClientAccountBanRes, err error) {
	request := g.RequestFromCtx(ctx)
	userId := gconv.Int(middlewares.ChatSupportAuth().GetIdentity(request.Context()))
	ticket, err := service.Ticket().HasTicket(ctx, req.ClientId)
	if ticket == nil || err != nil || ticket.ChatSupportId != userId {
		return nil, consts.ChatSupportNoAuth
	}
	_, err = service.Client().SetAccountStatus(ctx, req.ClientId, req.Status)
	return res, err
}
