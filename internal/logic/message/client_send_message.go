package message

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"cloudCustomerService/internal/model/entity"
	"cloudCustomerService/internal/service"
	"context"
)

type sMessage struct{}

func init() {
	service.RegisterMessage(New())

}

func New() *sMessage {
	return &sMessage{}
}

func (s *sMessage) ClientSendMessage(ctx context.Context, in *model.MessageAddInput) (int64, error) {

	message := &entity.Messages{
		TicketId:      in.TicketId,
		SenderType:    in.SenderType,
		ClientId:      in.ClientId,
		ChatSupportId: in.ChatSupportId,
		Content:       in.Content,
	}
	result, _ := dao.Messages.Ctx(ctx).Data(message).Insert()
	return result.LastInsertId()
}
