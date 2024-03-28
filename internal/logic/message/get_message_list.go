package message

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"context"
	"fmt"
)

func (s *sMessage) GetMessageList(ctx context.Context, in *model.MessageListInput) (*model.ClientMessageListOutput, error) {

	out := &model.ClientMessageListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	m := dao.Messages.Ctx(ctx).Page(in.Page, in.Size).OrderDesc(dao.Messages.Columns().Id).Where(dao.Messages.Columns().TicketId, in.TicketId)
	if !in.LastMessageId.IsEmpty() {
		m = m.Where(fmt.Sprintf("%s > %d", dao.Messages.Columns().Id, in.LastMessageId.Int()))
	}
	err := m.ScanAndCount(&out.List, &out.Total, false)
	return out, err

}
