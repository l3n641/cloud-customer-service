package message

import (
	"cloudCustomerService/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sMessage) UpdateMessageReadStatus(ctx context.Context, ticketId int, senderType int, msgIds []int) (int64, error) {

	result, _ := dao.Messages.Ctx(ctx).Where(g.Map{
		dao.Messages.Columns().TicketId:   ticketId,
		dao.Messages.Columns().SenderType: senderType,
		dao.Messages.Columns().IsRead:     0,
	}).WhereIn(dao.Messages.Columns().Id, msgIds).Data(g.Map{
		dao.Messages.Columns().IsRead:   1,
		dao.Messages.Columns().ReadTime: gtime.Now(),
	}).Update()

	return result.RowsAffected()
}
