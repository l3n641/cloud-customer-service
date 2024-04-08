// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/model/entity"
	"context"
)

type (
	ITicket interface {
		CreateTicket(ctx context.Context, clientId int, chatSupportId int, account string) (int, error)
		GetDetail(ctx context.Context, id int) (*model.TicketDetail, error)
		GetById(ctx context.Context, id int) (*entity.Tickets, error)
		HasTicket(ctx context.Context, clientId int) (*entity.Tickets, error)
		ChatSupportHasTicket(ctx context.Context, chatSupportId, ticketId int) (*entity.Tickets, error)
		SearchTicket(ctx context.Context, in *model.TicketSearchInput) (*model.TicketSearchOutput, error)
		DecrementCsUnreadQuantity(ctx context.Context, ticketId int, quantity int64) (int64, error)
		UpdateRemark(ctx context.Context, ticketId int, remark string) (int64, error)
		UpdateTicketByClient(ctx context.Context, ticketId int) (int64, error)
		UpdateTicketByChatSupport(ctx context.Context, ticketId int) (int64, error)
		UpdateTicketLastMessageTimeByChatSupport(ctx context.Context, ticketId int) (int64, error)
	}
)

var (
	localTicket ITicket
)

func Ticket() ITicket {
	if localTicket == nil {
		panic("implement not found for interface ITicket, forgot register?")
	}
	return localTicket
}

func RegisterTicket(i ITicket) {
	localTicket = i
}
