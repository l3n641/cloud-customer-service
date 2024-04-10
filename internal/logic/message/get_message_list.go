package message

import (
	"cloudCustomerService/internal/consts"
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sMessage) GetMessageList(ctx context.Context, in *model.MessageListInput) (*model.ClientMessageListOutput, error) {

	out := &model.ClientMessageListOutput{
		Page: in.Page,
		Size: in.Size,
		List: make([]model.ClientMessageItem, 0),
	}
	m := dao.Messages.Ctx(ctx).Page(in.Page, in.Size).OrderDesc(dao.Messages.Columns().Id).Where(dao.Messages.Columns().TicketId, in.TicketId)
	if !in.LastMessageId.IsEmpty() {
		m = m.Where(fmt.Sprintf("%s < %d", dao.Messages.Columns().Id, in.LastMessageId.Int()))
	}
	err := m.ScanAndCount(&out.List, &out.Total, false)
	return out, err

}

func (s *sMessage) ClientGetMessageList(ctx context.Context, in *model.ClientMessageListInput) (*model.ClientMessageListOutput, error) {

	out := &model.ClientMessageListOutput{
		Page: in.Page,
		Size: in.Size,
		List: make([]model.ClientMessageItem, 0),
	}
	m := dao.Messages.Ctx(ctx).Page(in.Page, in.Size).OrderDesc(dao.Messages.Columns().Id).Where(dao.Messages.Columns().TicketId, in.TicketId)
	if !in.LastMessageId.IsEmpty() {
		m = m.Where(fmt.Sprintf("%s < %d", dao.Messages.Columns().Id, in.LastMessageId.Int()))
	}
	if !in.OnlyUnread.IsEmpty() {
		m = m.Where(g.Map{dao.Messages.Columns().SenderType: consts.SendMessageToClient, dao.Messages.Columns().IsRead: 0})
	}
	err := m.ScanAndCount(&out.List, &out.Total, false)
	return out, err

}
