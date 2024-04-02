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

func (s *sMessage) SendMessage(ctx context.Context, in *model.MessageAddInput) (*entity.Messages, error) {

	message := &entity.Messages{
		TicketId:      in.TicketId,
		SenderType:    in.SenderType,
		ClientId:      in.ClientId,
		ChatSupportId: in.ChatSupportId,
		Content:       in.Content,
		RefererUrl:    in.RefererUrl,
		MsgType:       in.MsgType,
	}
	result, err := dao.Messages.Ctx(ctx).Data().Insert(message)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return s.GetMessageById(ctx, id)

}
