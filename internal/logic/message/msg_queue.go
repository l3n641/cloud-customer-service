package message

import (
	"cloudCustomerService/internal/consts"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"strconv"
)

func (s *sMessage) PushChatSupportMessage(ctx context.Context, messages *entity.Messages) (int64, error) {

	redisClient := g.Redis(consts.RedisMsgQueueName)
	key := consts.ChatSupportQueueKey + strconv.Itoa(messages.ChatSupportId)
	return redisClient.LPush(ctx, key, &model.ChatSupportMessageItem{
		Id:            messages.Id,
		CreateAt:      messages.CreateAt,
		TicketId:      messages.TicketId,
		SenderType:    messages.SenderType,
		ClientId:      messages.ClientId,
		ChatSupportId: messages.ChatSupportId,
		MsgType:       messages.MsgType,
		Content:       messages.Content,
		IsRead:        messages.IsRead,
		ReadTime:      messages.ReadTime,
		RefererUrl:    messages.RefererUrl,
	})

}
