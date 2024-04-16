package messageTemplate

import (
	"cloudCustomerService/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sMessageTemplate) DeleteMessage(ctx context.Context, id int) (int64, error) {
	result, _ := dao.MessageTemplates.Ctx(ctx).Where(g.Map{
		dao.MessageTemplates.Columns().Id: id,
	}).Delete()

	return result.RowsAffected()

}
