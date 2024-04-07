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

func (c *ControllerTicket) DetailTicket(ctx context.Context, req *ticket.DetailTicketReq) (res *ticket.DetailTicketRes, err error) {
	request := g.RequestFromCtx(ctx)
	userId := gconv.Int(middlewares.ChatSupportAuth().GetIdentity(request.Context()))
	ticketData, err := service.Ticket().ChatSupportHasTicket(ctx, userId, req.Id)
	if ticketData == nil || err != nil {
		return nil, consts.ChatSupportNoAuth
	}
	data, err := service.Ticket().GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	clientInfo := ticket.ClientInfo{
		Id:                 data.Client.Id,
		Email:              data.Client.Email,
		Phone:              data.Client.Phone,
		Domain:             data.Client.Domain,
		MerchantId:         data.Client.MerchantId,
		Ip:                 data.Client.Ip,
		BrowserFingerprint: data.Client.BrowserFingerprint,
		Area:               data.Client.Area,
		UserAgent:          data.Client.UserAgent,
		Lang:               data.Client.Lang,
	}

	ticketInfo := ticket.DetailTicketInfo{
		Id:              data.Ticket.Id,
		CreateAt:        data.Ticket.CreateAt,
		LastMessageTime: data.Ticket.LastMessageTime,
		Remark:          data.Ticket.Remark,
	}

	return &ticket.DetailTicketRes{
		ClientInfo:       clientInfo,
		DetailTicketInfo: ticketInfo,
	}, err

}
