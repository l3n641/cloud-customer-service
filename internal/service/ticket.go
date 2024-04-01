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
		CreateTicket(ctx context.Context, clientId int, chatSupportId int) (int, error)
		HasTicket(ctx context.Context, clientId int) (*entity.Tickets, error)
		SearchTicket(ctx context.Context, in *model.TicketSearchInput) (*model.TicketSearchOutput, error)
		UpdateTicketByClient(ctx context.Context, ticketId int) (int64, error)
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
