package ticket

import (
	"cloudCustomerService/internal/dao"
	"context"
)

func (s *sTicket) DecrementCsUnreadQuantity(ctx context.Context, ticketId int, quantity int64) (int64, error) {

	result, _ := dao.Tickets.Ctx(ctx).Where(dao.Tickets.Columns().Id, ticketId).Decrement(dao.Tickets.Columns().CsUnreadMsgCount, quantity)
	return result.RowsAffected()
}
