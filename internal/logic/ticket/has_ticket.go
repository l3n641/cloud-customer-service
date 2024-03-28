package ticket

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTicket) HasTicket(ctx context.Context, clientId int) (*entity.Tickets, error) {
	var ticket *entity.Tickets
	err := dao.Tickets.Ctx(ctx).Where(g.Map{
		dao.Tickets.Columns().ClientId: clientId,
	}).Scan(&ticket)

	return ticket, err

}
