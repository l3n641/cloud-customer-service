package message

import (
	"cloudCustomerService/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sMessage) GetUnreadMessageQuantity(ctx context.Context, senderType, ticketId int) (int, error) {

	return dao.Messages.Ctx(ctx).Where(g.Map{
		dao.Messages.Columns().SenderType: senderType,
		dao.Messages.Columns().TicketId:   ticketId,
		dao.Messages.Columns().IsRead:     0,
	}).Count()
}
