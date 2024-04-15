package messageTemplate

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"context"
)

func (s *sMessageTemplate) GetMessageList(ctx context.Context, in *model.MessageTemplateListInput) (*model.MessageTemplateListOutput, error) {

	out := &model.MessageTemplateListOutput{
		List: make([]model.MessageTemplateItem, 0),
	}
	m := dao.MessageTemplates.Ctx(ctx).Page(in.Page, in.Size).OrderDesc(dao.MessageTemplates.Columns().Id)

	if !in.MsgType.IsEmpty() {
		m = m.Where(dao.MessageTemplates.Columns().MsgType, in.MsgType.Int())
	}

	if in.Lang != "" {
		m = m.Where(dao.MessageTemplates.Columns().Lang, in.Lang)
	}

	if !in.MsgType.IsEmpty() {
		m = m.Where(dao.MessageTemplates.Columns().MsgType, in.MsgType.Int())
	}
	err := m.ScanAndCount(&out.List, &out.Total, false)
	return out, err

}
