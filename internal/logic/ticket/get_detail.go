package ticket

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/model/entity"
	"context"
)

func (s *sTicket) GetDetail(ctx context.Context, id int) (*model.TicketDetail, error) {

	detail := model.TicketDetail{}
	if err := dao.Tickets.Ctx(ctx).Where(dao.Tickets.Columns().Id, id).Scan(&detail.Ticket); err != nil {
		return nil, err
	}

	err := dao.Clients.Ctx(ctx).Where(dao.Clients.Columns().Id, detail.Ticket.ClientId).Scan(&detail.Client)
	return &detail, err
}

func (s *sTicket) GetById(ctx context.Context, id int) (*entity.Tickets, error) {

	var data *entity.Tickets
	err := dao.Tickets.Ctx(ctx).Where(dao.Tickets.Columns().Id, id).Scan(&data)
	return data, err
}
