package message

import (
	"cloudCustomerService/internal/consts"
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"strconv"
)

func (s *sMessage) ReadMessage(ctx context.Context, ticketId, senderType int) (int64, error) {

	result, _ := dao.Messages.Ctx(ctx).Where(g.Map{
		dao.Messages.Columns().TicketId:   ticketId,
		dao.Messages.Columns().SenderType: senderType,
		dao.Messages.Columns().IsRead:     0,
	}).Data(g.Map{
		dao.Messages.Columns().IsRead:   1,
		dao.Messages.Columns().ReadTime: gtime.Now(),
	}).Update()

	return result.RowsAffected()
}

func (s *sMessage) ClientReadAllMessage(ctx context.Context, ticketId, ChatSupportId int) (row int64, err error) {

	row, err = s.ReadMessage(ctx, ticketId, consts.SendMessageToClient)

	if row > 0 {
		redisClient := g.Redis(consts.RedisMsgQueueName)
		key := consts.TicketQueueKey + strconv.Itoa(ChatSupportId)
		redisClient.LPush(ctx, key, model.ClientReadTicket{TicketId: ticketId, ReadTime: gtime.Now()})
	}
	return row, err
}
