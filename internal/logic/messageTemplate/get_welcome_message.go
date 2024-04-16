package messageTemplate

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sMessageTemplate) GetWelcomeMessage(ctx context.Context) (*[]model.MessageTemplateItem, error) {
	var data = make([]model.MessageTemplateItem, 0)
	err := dao.MessageTemplates.Ctx(ctx).Where(g.Map{
		dao.MessageTemplates.Columns().MsgType: 1,
	}).Scan(&data)

	return &data, err

}
