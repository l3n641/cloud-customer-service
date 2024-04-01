package ticket

import (
	"cloudCustomerService/internal/dao"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sTicket) UpdateTicketByClient(ctx context.Context, ticketId int) (int64, error) {

	result, _ := dao.Tickets.Ctx(ctx).Data(g.Map{
		dao.Tickets.Columns().LastMessageTime: gtime.Now(),
		dao.Tickets.Columns().CsUnreadMsgCount: &gdb.Counter{
			Field: dao.Tickets.Columns().CsUnreadMsgCount,
			Value: 1,
		},
	}).Where(dao.Tickets.Columns().Id, ticketId).Update()
	return result.RowsAffected()
}
