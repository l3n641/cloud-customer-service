package ticket

import (
	"cloudCustomerService/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sTicket) UpdateRemark(ctx context.Context, ticketId int, remark string) (int64, error) {

	result, _ := dao.Tickets.Ctx(ctx).Data(g.Map{
		dao.Tickets.Columns().Remark: remark,
	}).Where(dao.Tickets.Columns().Id, ticketId).Update()
	return result.RowsAffected()
}
