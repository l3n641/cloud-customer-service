package messageTemplate

import (
	"cloudCustomerService/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sMessageTemplate) UpdateMessageTemplate(ctx context.Context, id, msgType int, lang, content string) (int64, error) {

	data := g.Map{
		dao.MessageTemplates.Columns().Content: content,
		dao.MessageTemplates.Columns().MsgType: msgType,
		dao.MessageTemplates.Columns().Lang:    lang,
	}
	result, err := dao.MessageTemplates.Ctx(ctx).Where(dao.MessageTemplates.Columns().Id, id).Data(data).Update()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()

}
